%option noyywrap
%option yylineno

%{
#include "grammar.tab.h"
#include <stdlib.h>
void yyerror(char *);
%}

%%
[1-9][0-9]*			          {yylval.ival=atoi(yytext); return INTVAL;}
(true)/[^a-zA-Z0-9_]		  {yylval.bval=TRUE; return TRUE_;}
(false)/[^a-zA-Z0-9_]		  {yylval.bval=FALSE; return FALSE_;}
(undefined)/[^a-zA-Z0-9_]	  {yylval.bval=UNDEFINED; return UNDEFINED_;}
(int)/[^a-zA-Z0-9_]		      {yylval.types_=INT; return INT_;}
(Sint)/[^a-zA-Z0-9_]		  {yylval.types_=SHORT; return SHORT_;}
(bool)/[^a-zA-Z0-9_]		  {yylval.types_=BOOL; return BOOL_;}
(vector)/[^a-zA-Z0-9_]		  {yylval.types_=VECTOR; return VECTOR_;}
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
(begin)/[^a-zA-Z0-9_]		  {return BEGIN_;}
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
(work)/[^a-zA-Z0-9_]		      {return WORK;}
[\[\]\|\{\},\(\);\n]          {return *yytext;}
[a-zA-Z_][a-zA-Z_0-9]*        {yylval.string=new std::string(yytext); return NAME;}

[ \t]				          {};
.				              {yyerror(yytext);}

%%
