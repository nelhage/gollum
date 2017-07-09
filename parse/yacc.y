%{

package parse

import (
   "nelhage.com/lambda"
)

%}

%union {
    ast lambda.AST
    asts []lambda.AST
    tok *tokenStruct
    toks []tokenStruct
}

%type   <ast>           expression literal condition
%type   <ast>           variable abstraction application

%type   <asts>          expressionlist
%type   <asts>          expressions

%type   <toks>          varlist
%type   <toks>          vars

%token  <tok>           tokFunc tokIf tokThen tokElse tokEnd
%token  <tok>           tokBoolean tokNumber tokStr
%token  <tok>           tokIdent
%token  <tok>           ',' '(' ')' '{' '}'

%nonassoc ')'

%%

program:
                expression
                {
                    yylex.(*lexer).result = $1
                }

expression:
                literal
        |       condition
        |       variable
        |       abstraction
        |       application
        |       '(' expression ')'
                {
                    $$ = $2
                }

literal:
                tokBoolean
                {
                    $$ = &lambda.Boolean{
                        Loc: $1.loc,
                        Value: $1.val.(string) == "true",
                    }
                }
        |       tokNumber
                {
                    $$ = &lambda.Integer{
                        Loc: $1.loc,
                        Value: $1.val.(int64),
                    }
                }
        |       tokStr
                {
                    $$ = &lambda.String{
                        Loc: $1.loc,
                        Value: $1.val.(string),
                    }
                }

condition:
                tokIf expression tokThen expression tokElse expression tokEnd
                {
                    $$ = &lambda.If{
                        Loc: extend($1.loc, $6.Location()),
                        Condition: $2,
                        Consequent: $4,
                        Alternate: $6,
                    }
                }

variable:       tokIdent
                {
                    $$ = &lambda.Variable{
                        Loc: $1.loc,
                        Var: $1.val.(string),
                    }
                }

abstraction:
                tokFunc '(' varlist ')' '{' expression '}'
                {
                    vars := []string{}
                    for _, tok := range $3 {
                        vars = append(vars, tok.val.(string))
                    }
                    $$ = &lambda.Abstraction {
                        Loc: extend($1.loc, $7.loc),
                        Vars: vars,
                        Body: $6,
                    }
                }

varlist:
                {
                    $$ = []tokenStruct{}
                }
        |       vars
        |       vars ','

vars:           tokIdent
                {
                    $$ = []tokenStruct{*$1}
                }
        |       vars ',' tokIdent
                {
                    $$ = append($1, *$3)
                }

application:
                expression '(' expressionlist ')'
                {
                    $$ = &lambda.Application{
                        Loc: extend($1.Location(), $4.loc),
                        Func: $1,
                        Args: $3,
                    }
                }

expressionlist:
                {
                    $$ = []lambda.AST{}
                }
        |       expressions
        |       expressions ','

expressions:
                expression
                {
                    $$ = []lambda.AST{$1}
                }
        |       expressions ',' expression
                {
                    $$ = append($1, $3)
                }

%%
