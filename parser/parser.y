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
%token ADD

%left ADD

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
| expr ADD expr
{
  $$ = createAddNode($1, $3)
}
%%
