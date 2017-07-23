%{

package parse

import (
   "github.com/nelhage/gollum"
)

var keywords = map[string]token{
	"if":   tokIf,
	"else": tokElse,
	"let":  tokLet,
	"rec": tokRec,
	"in":   tokIn,

	"fn":    tokFunc,
	"true":  tokBoolean,
	"false": tokBoolean,
}

%}

%union {
    ast gollum.AST
    asts []gollum.AST
    tok *tokenStruct
}

%type   <ast>           program
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

%token  <tok>           tokFunc tokIf tokElse tokLet tokRec tokIn
%token  <tok>           tokBoolean tokNumber tokStr
%token  <tok>           tokIdent
%token  <tok>           tokArrow
%token  <tok>           tokIntType
%token  <tok>           ',' '(' ')' '{' '}' ':'

%right tokArrow
%nonassoc ','
%nonassoc ')'

%%

top:            program
                {
                    yylex.(*lexer).expression = $1
                }
        |       tokIntType type
                {
                    yylex.(*lexer).ty = $2
                }

program:        expression

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
                    $$ = &gollum.Boolean{
                        Loc: $1.loc,
                        Value: $1.val.(string) == "true",
                    }
                }
        |       tokNumber
                {
                    $$ = &gollum.Integer{
                        Loc: $1.loc,
                        Value: $1.val.(int64),
                    }
                }
        |       tokStr
                {
                    $$ = &gollum.String{
                        Loc: $1.loc,
                        Value: $1.val.(string),
                    }
                }

condition:
                tokIf expression brackExpr tokElse brackExpr
                {
                    $$ = &gollum.If{
                        Loc: extend($1.loc, $5.Location()),
                        Condition: $2,
                        Consequent: $3,
                        Alternate: $5,
                    }
                }

variable:       tokIdent
                {
                    $$ = &gollum.Variable{
                        Loc: $1.loc,
                        Var: $1.val.(string),
                    }
                }

abstraction:
                tokFunc '(' varList ')' brackExpr
                {
                    $$ = &gollum.Abstraction {
                        Loc: extend($1.loc, $5.Location()),
                        Vars: $3,
                        Body: $5,
                    }
                }

varList:
                {
                    $$ = []gollum.AST{}
                }
        |       vars
        |       vars ','

vars:           typedecl
                {
                    $$ = []gollum.AST{$1}
                }
        |       vars ',' typedecl
                {
                    $$ = append($1, $3)
                }

typedecl:
                tokIdent
                {
                    $$ = &gollum.TypedName{
                        Loc: $1.loc,
                        Name: $1.val.(string),
                        Type: nil,
                    }
                }
        |       tokIdent ':' type
                {
                    $$ = &gollum.TypedName{
                        Loc: $1.loc,
                        Name: $1.val.(string),
                        Type: $3,
                    }
                }

type:
                tokIdent
                {
                    $$ = &gollum.TyName{
                        Loc: $1.loc,
                        Type: $1.val.(string),
                    }
                }
        |       type tokArrow type
                {
                    $$ = &gollum.TyArrow{
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
                    $$ = &gollum.TyTuple{
                        Loc: extend($1.loc, $3.loc),
                        Elts: $2,
                    }
                }
        |       '(' tupleType ',' ')'
                {
                    $$ = &gollum.TyTuple{
                        Loc: extend($1.loc, $3.loc),
                        Elts: $2,
                    }
                }

tupleType:
                type %prec ','
                {
                    $$ = []gollum.AST{$1}
                }
        |       tupleType ',' type
                {
                    $$ = append($1, $3)
                }

application:
                expression '(' expressionList ')'
                {
                    $$ = &gollum.Application{
                        Loc: extend($1.Location(), $4.loc),
                        Func: $1,
                        Args: $3,
                    }
                }

expressionList:
                {
                    $$ = []gollum.AST{}
                }
        |       expressions
        |       expressions ','

expressions:
                expression
                {
                    $$ = []gollum.AST{$1}
                }
        |       expressions ',' expression
                {
                    $$ = append($1, $3)
                }

let:            tokLet bindingList tokIn brackExpr
                {
                    $$ = &gollum.Let{
                        Loc: extend($1.loc, $4.Location()),
                        Bindings: $2,
                        Body: $4,
                    }
                }
        |       tokLet tokRec bindingList tokIn brackExpr
                {
                    $$ = &gollum.Let{
                        Loc: extend($1.loc, $5.Location()),
                        Bindings: $3,
                        Body: $5,
                        Recursive: true,
                    }
                }

bindingList:
                {
                    $$ = []gollum.AST{}
                }
        |       bindings
        |       bindings ','

bindings:
                binding
                {
                    $$ = []gollum.AST{$1}
                }
        |       bindings ',' binding
                {
                    $$ = append($1, $3)
                }

binding:
                typedecl '=' expression
                {
                    $$ = &gollum.NameBinding{
                        Loc: extend($1.Location(), $3.Location()),
                        Var: $1,
                        Value: $3,
                    }
                }

%%
