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
%token BIN_OP_1
%token BIN_OP_2
%token ASSIGNMENT
%token END
%token IF
%token FOR
%token WHILE
%token FUNCTION
%token RETURN

%token LP
%token RP
%token LB
%token RB

%right ASSIGNMENT
%left BIN_OP_1
%left BIN_OP_2
%left LP
%left RP

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
| FOR LP statement statement expr RP LB statements RB
{
  $$ = createForNode($3, $4, $5, $8)
}
| WHILE LP expr RP LB statements RB
{
  $$ = createWhileNode($3, $6)
}
| RETURN expr END
{
  $$ = createReturnNode($2)
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
| expr BIN_OP_2 expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr BIN_OP_1 expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr ASSIGNMENT expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| FUNCTION LP RP LB statements RB
{
  $$ = createFunctionNode($5)
}
| expr LP RP
{
  $$ = createCallNode($1)
}
%%
