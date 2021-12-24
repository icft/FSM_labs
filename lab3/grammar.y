%{
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <structs.h>
#include <lex.yy.c>

extern FILE *yyin;

Node *opr(int oper, int num, ...);
Node *id(std::string s);
Node *con(int value);
void freeNode(Node *n);
int exec(Node *p);
int yylex(void);
void init(void);
void yyerror(chat *s);

std::map<std::string, int> sym;
%}

%union {
	int ival;
	bool bval;
	std::string str;
	nodeType* ptr;
};

%token <ival> INTVAL
%token <bval> TRUE FALSE UNDEFINED
%token <str> NAME
%token INT SHORT BOOL VECTOR OF BEGIN DO WHILE IF THEN MOVE SET RIGHT LEFT LMS FUNC RETURN
%nonassoc IFX
%nonassoc ELSE

%left SMALLER LARGER
%left ADD SUB NOT OR AND

%type <ptr> stmt expr stmt_list function type params names directions fdecl

%%

program:
      '\n'                                                  {}
    | program stmt                                          {}
    | program stmt '\n'                                     {}
    ;

stmt:
      ';'                                                   {}
    | expr ';'                                              {}
    | arg  ';'                                              {}
    | arg SET expr ';'                                      {}
    | NAME SET expr ';'                                     {}
    | NAME operands ';'                                     {}
    | SIZEOF '(' type ')' ';'                               {}
    | SIZEOF '(' NAME ')' ';'                               {}
    | DO stmt WHILE expr                                    {}
    | IF expr THEN stmt %prec IFX                           {}
    | IF expr THEN stmt ELSE stmt                           {}
    | directions ';'                                        {}
    | BEGIN stmt END ';'                                    {}
    | BEGIN END                                             {}
    | fdecl                                                 {}
    ;

fdecl:
    FUNC NAME '(' arglist ')' BEGIN stmt_list RETURN expr   {}

arglist:
	  '(' ')'                                               {}
	| '(' args ')'                                          {}
	;

args:
	  arg                                                   {}
	| args ',' arg                                          {}
	;

arg:
      type NAME                                             {}
    ;

names:
      NAME                                                  {}
    | names NAME                                            {}
    ;

stmt_list:
      stmt                                                  {}
    | stmt_list stmt                                        {}
    ;

expr:
      INTVAL                                                {}
    | NAME                                                  {}
    | TRUE                                                  {}
    | FALSE                                                 {}
    | UNDEFINED                                             {}
    | expr ADD expr                                         {}
    | expr SUB expr                                         {}
    | expr '|' expr SMALLER                                 {}
    | expr '|' expr LARGER                                  {}
    | expr OR expr                                          {}
    | expr AND expr                                         {}
    | expr NOT OR expr                                      {}
    | expr NOT AND expr                                     {}
    | '(' expr ')'                                          {}
    ;

vecdecl:
       vecof type NAME '[' expr ']'                         {}
    |  vecof type NAME '[' expr ']'
    |  vecof type NAME SET '{' vecoperands '}'              {}
    ;

vecoperands:
      vecoperandslist                                       {}
    | '{' vecoperands '}' ',' vecoperands                   {}
    |
    ;

vecoperandslist:
      expr                                                  {}
    | vecoperandslist ',' expr                              {}
    ;

vecof:
      VECTOR OF                                             {}
    | vecof VECTOR OF                                       {}
    ;

type:
      INT                                                   {}
    | SHORT                                                 {}
    | BOOL                                                  {}
    ;

directions:
      MOVE RIGHT                                            {}
    | MOVE LEFT                                             {}
    | MOVE                                                  {}
    | RIGHT                                                 {}
    | LEFT                                                  {}
    ;

operands:
      '(' ')'                                               {}
    | '(' oplist ')'                                        {}
    ;

oplist:
     expr                                                   {}
   | oplist ' ' expr                                        {}
   ;
%%
