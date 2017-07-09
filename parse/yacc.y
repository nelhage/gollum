%{

package parse

import (
   "nelhage.com/lambda"
)

%}

%union {
    ast *lambda.AST
    tok *tokenStruct
}

%%

program:

%%
