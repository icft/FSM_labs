%option noyywrap

%{
#include "grammar.tab.h"
#include <stdlib.h>
void yyerror(char *);
%}

%%
[1-9][0-9]*			          {yyval.ival=atoi(yytext); return INTVAL;}
(true)/[^a-zA-Z0-9_]		  {yyval.bval = 1; return TRUE;}
(false)/[^a-zA-Z0-9_]		  {yyval.bval = -1; return FALSE;}
(undefined)/[^a-zA-Z0-9_]	  {yyval.bval = 0; return UNDEFINED;}
(int)/[^a-zA-Z0-9_]		      {return INT;}
(Sint)/[^a-zA-Z0-9_]		  {return SHORT;}
(bool)/[^a-zA-Z0-9_]		  {return BOOL;}
(vector)/[^a-zA-Z0-9_]		  {return VECTOR;}
(of)/[^a-zA-Z0-9_]		      {return OF;}
(sizeof)/[^a-zA-Z0-9_]		  {return SIZEOF;}
(set)/[^a-zA-Z0-9_]		      {return SET;}
(add)/[^a-zA-Z0-9_]		      {return ADD;}
(sub)/[^a-zA-Z0-9_]		      {return SUB;}
(smaller)/[^a-zA-Z0-9_]		  {return SMALLER;}
(larger)/[^a-zA-Z0-9_]		  {return LARGER;}
(not)/[^a-zA-Z0-9_]		      {return NOT;}
(or)/[^a-zA-Z0-9_]		      {return OR;}
(and)/[^a-zA-Z0-9_]		      {return AND;}
(begin)/[^a-zA-Z0-9_]		  {return BEGIN;}
(end)/[^a-zA-Z0-9_]		      {return END;}
(do)/[^a-zA-Z0-9_]		      {return DO;}
(while)/[^a-zA-Z0-9_]		  {return WHILE;}
(if)/[^a-zA-Z0-9_]		      {return IF;}
(then)/[^a-zA-Z0-9_]		  {return THEN;}
(else)/[^a-zA-Z0-9_]		  {return ELSE;}
(move)/[^a-zA-Z0-9_]		  {return MOVE;}
(left)/[^a-zA-Z0-9_]		  {return LEFT;}
(right)/[^a-zA-Z0-9_]		  {return RIGHT;}
(lms)/[^a-zA-Z0-9_]		      {return LMS;}
(function)/[^a-zA-Z0-9_]	  {return FUNC;}
(return)/[^a-zA-Z0-9_]		  {return RETURN;}
[\[\]\|\{\},\(\);\n]          {return *yytext;}
[a-zA-Z_][a-zA-Z_0-9]*        {yylval.string=yytext;
                               return NAME;}

[ \t]				          {};
.				              yyerror(text);

%%
