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
%token IDENTIFIER
%token BINARY_OPERATOR
%token ASSIGNMENT
%token END

%right ASSIGNMENT
%left BINARY_OPERATOR

%%
program: statements
{
  setParseResult(yylex, $1)
}

statements: expr END
{
  $$ = appendStatement(nil, $1)
}
| statements expr END
{
  $$ = appendStatement(&$1, $2)
}

expr: NUMBER
{
  $$ = createNumberNode($1)
}
| STRING
{
  $$ = createStringNode($1)
}
| IDENTIFIER
{
  $$ = createIdentifierNode($1)
}
| expr BINARY_OPERATOR expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr ASSIGNMENT expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
%%
