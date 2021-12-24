%{
#include "lex.yy.c"
#include "Nodes.h"
int yylex(void);
void yyerror(chat *s);
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
%token INT SHORT BOOL VECTOR OF BEGIN_ DO WHILE IF THEN MOVE SET RIGHT LEFT LMS FUNC RETURN SIZEOF 
%nonassoc IFX
%nonassoc ELSE
%nonassoc END
%left SMALLER LARGER
%left ADD SUB NOT OR AND

%type <Node> stmt expr stmt_list type names directions fdecl fcall

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
    | BEGIN_ stmt END                                       {}
    | BEGIN_ END                                            {}
    | fdecl                                                 {}
    | fcall                                                 {}
    ;

fdecl:
    FUNC NAME '(' arglist ')' BEGIN_ stmt_list RETURN expr  {}

fcall:
    NAME '(' names ')' ';'                                 {}

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
      INTVAL                                                {$$=$1;}
    | NAME                                                  {$$=$1;}
    | TRUE                                                  {$$=$1;}
    | FALSE                                                 {$$=$1;}
    | UNDEFINED                                             {$$=$1;}
    | expr ADD expr                                         {$$=$1+$3;}
    | expr SUB expr                                         {$$=$1-$3;}
    | expr '|' expr SMALLER                                 {if ($1<$3) {printf("true");}
                                                             if ($1>$3) {printf("false");}
                                                             if ($1==$3) {printf("undefined");}}
    | expr '|' expr LARGER                                  {if ($1>$3) {printf("true");}
                                                             if ($1<$3) {printf("false");}
                                                             if ($1==$3) {printf("undefined");}}
    | expr OR expr                                          {$$=$1 || $3;}
    | expr AND expr                                         {$$=$1 && $3;}
    | expr NOT OR expr                                      {}
    | expr NOT AND expr                                     {}
    | '(' expr ')'                                          {}
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

void yyerror(char *s) {
   std::string str;
   str = c;
   std::cout << str << std::endl;
}

int main(void) {
   yyparse();
   return 0;
}
	
