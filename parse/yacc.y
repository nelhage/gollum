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
}

%type   <ast>           expression brackExpr literal condition
%type   <ast>           variable abstraction application let

%type   <asts>          expressionList
%type   <asts>          expressions

%type   <asts>          varList
%type   <asts>          vars

%type   <ast>           typedecl
%type   <ast>           type
%type   <asts>          tupleType

%type   <asts>          bindingList
%type   <asts>          bindings
%type   <ast>           binding

%token  <tok>           tokFunc tokIf tokElse tokLet tokIn
%token  <tok>           tokBoolean tokNumber tokStr
%token  <tok>           tokIdent
%token  <tok>           tokArrow
%token  <tok>           ',' '(' ')' '{' '}' ':'

%right tokArrow
%nonassoc ','
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
        |       brackExpr
        |       let
        |       '(' expression ')'
                {
                    $$ = $2
                }

brackExpr:
                '{'expression '}'
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
                tokIf expression brackExpr tokElse brackExpr
                {
                    $$ = &lambda.If{
                        Loc: extend($1.loc, $5.Location()),
                        Condition: $2,
                        Consequent: $3,
                        Alternate: $5,
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
                tokFunc '(' varList ')' brackExpr
                {
                    $$ = &lambda.Abstraction {
                        Loc: extend($1.loc, $5.Location()),
                        Vars: $3,
                        Body: $5,
                    }
                }

varList:
                {
                    $$ = []lambda.AST{}
                }
        |       vars
        |       vars ','

vars:           typedecl
                {
                    $$ = []lambda.AST{$1}
                }
        |       vars ',' typedecl
                {
                    $$ = append($1, $3)
                }

typedecl:
                tokIdent
                {
                    $$ = &lambda.TypedName{
                        Loc: $1.loc,
                        Name: $1.val.(string),
                        Type: nil,
                    }
                }
        |       tokIdent ':' type
                {
                    $$ = &lambda.TypedName{
                        Loc: $1.loc,
                        Name: $1.val.(string),
                        Type: $3,
                    }
                }

type:
                tokIdent
                {
                    $$ = &lambda.TyName{
                        Loc: $1.loc,
                        Type: $1.val.(string),
                    }
                }
        |       type tokArrow type
                {
                    $$ = &lambda.TyArrow{
                        Loc: extend($1.Location(), $3.Location()),
                        Dom: $1,
                        Range: $3,
                    }
                }
        |       '(' type ')'
                {
                    $$ = $2
                }
        |       '(' tupleType ')'
                {
                    $$ = &lambda.TyTuple{
                        Loc: extend($1.loc, $3.loc),
                        Elts: $2,
                    }
                }
        |       '(' tupleType ',' ')'
                {
                    $$ = &lambda.TyTuple{
                        Loc: extend($1.loc, $3.loc),
                        Elts: $2,
                    }
                }

tupleType:
                type %prec ','
                {
                    $$ = []lambda.AST{$1}
                }
        |       tupleType ',' type
                {
                    $$ = append($1, $3)
                }

application:
                expression '(' expressionList ')'
                {
                    $$ = &lambda.Application{
                        Loc: extend($1.Location(), $4.loc),
                        Func: $1,
                        Args: $3,
                    }
                }

expressionList:
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

let:            tokLet bindingList tokIn brackExpr
                {
                    $$ = &lambda.Let{
                        Loc: extend($1.loc, $4.Location()),
                        Bindings: $2,
                        Body: $4,
                    }
                }

bindingList:
                {
                    $$ = []lambda.AST{}
                }
        |       bindings
        |       bindings ','

bindings:
                binding
                {
                    $$ = []lambda.AST{$1}
                }
        |       bindings ',' binding
                {
                    $$ = append($1, $3)
                }

binding:
                typedecl '=' expression
                {
                    $$ = &lambda.NameBinding{
                        Loc: extend($1.Location(), $3.Location()),
                        Var: $1,
                        Value: $3,
                    }
                }

%%
