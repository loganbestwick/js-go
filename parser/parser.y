%{
package parser

import "github.com/loganbestwick/js-go/syntax"
%}

%union {
  s string
  node syntax.Node
}

%token NUMBER
%token ADD

%left ADD

%%
program: expr
{
  setParseResult(yylex, $1)
}

expr: NUMBER
{
  $$ = createValueNode($1)
}
| expr ADD expr
{
  $$ = createAddNode($1, $3)
}
%%
