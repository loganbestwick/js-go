%{
package parser

import "github.com/loganbestwick/js-go/syntax"
%}

%union {
  s string
  node syntax.Node
}

%token BOOLEAN
%token NUMBER
%token STRING
%token IDENTIFIER
%token BINARY_OPERATOR
%token ASSIGNMENT
%token END
%token IF
%token LP
%token RP
%token LB
%token RB

%right ASSIGNMENT
%left BINARY_OPERATOR

%%
program: statements
{
  setParseResult(yylex, $1)
}

statements: statement
{
  $$ = appendStatement(nil, $1)
}
| statements statement
{
  $$ = appendStatement(&$1, $2)
}

statement: expr END
{
  $$ = $1
}
| IF LP expr RP LB statements RB
{
  $$ = createIfNode($3, $6)
}

expr: BOOLEAN
{
  $$ = createBooleanNode($1)
} | NUMBER
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
