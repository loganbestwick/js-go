%{
package parser

import "github.com/loganbestwick/js-go/syntax"
%}

%union {
  s string
  node syntax.Node
}

%token NUMBER
%token STRING
%token BINARY_OPERATOR

%left BINARY_OPERATOR

%%
program: expr
{
  setParseResult(yylex, $1)
}

expr: NUMBER
{
  $$ = createNumberNode($1)
}
| STRING
{
  $$ = createStringNode($1)
}
| expr BINARY_OPERATOR expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
%%
