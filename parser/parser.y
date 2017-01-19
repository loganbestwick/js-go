%{
package parser
%}

%union {
  s string
  node Node
}

%token NUMBER

%%
program: expr
{
  setParseResult(yylex, $1)
}

expr: NUMBER
{
  $$ = node($1)
}
%%