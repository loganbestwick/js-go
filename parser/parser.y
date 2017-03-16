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
%token END
%token ASSIGN
%token BINARY_OPERATOR

%right ASSIGN
%left BINARY_OPERATOR

%%
program: statements
{
  setParseResult(yylex, $1)
}

statements: statement
{
  $$ = createStatementsNode($1)
}
| statements statement
{
  $$ = appendStatementsNode($1, $2)
}

statement: expr END
{
  $$ = $1
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
  $$ = createVariableNode($1)
}
| expr ASSIGN expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr BINARY_OPERATOR expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
%%
