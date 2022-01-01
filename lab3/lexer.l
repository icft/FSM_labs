%option noyywrap
%option yylineno

%{
#include "grammar.tab.h"
#include <stdlib.h>
void yyerror(char *);
%}

%%
[1-9][0-9]*			          {yylval.ival=atoi(yytext); yylloc.last_line=yylloc.first_line =yylineno; return INTVAL;}
(true)/[^a-zA-Z0-9_]		  {yylval.bval = 1; yylloc.last_line=yylloc.first_line =yylineno; return TRUE;}
(false)/[^a-zA-Z0-9_]		  {yylval.bval = -1; yylloc.last_line=yylloc.first_line =yylineno; return FALSE;}
(undefined)/[^a-zA-Z0-9_]	  {yylval.bval = 0; yylloc.last_line=yylloc.first_line =yylineno; return UNDEFINED;}
(int)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return INT;}
(Sint)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return SHORT;}
(bool)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return BOOL;}
(vector)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return VECTOR;}
(of)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return OF;}
(sizeof)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return SIZEOF;}
(set)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return SET;}
(add)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return ADD;}
(sub)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return SUB;}
(smaller)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return SMALLER;}
(larger)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return LARGER;}
(not)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return NOT;}
(or)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return OR;}
(and)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return AND;}
(begin)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return BEGIN;}
(end)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return END;}
(do)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return DO;}
(while)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return WHILE;}
(if)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return IF;}
(then)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return THEN;}
(else)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return ELSE;}
(move)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return MOVE;}
(left)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return LEFT;}
(right)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return RIGHT;}
(lms)/[^a-zA-Z0-9_]		      {yylloc.last_line=yylloc.first_line =yylineno; return LMS;}
(function)/[^a-zA-Z0-9_]	  {yylloc.last_line=yylloc.first_line =yylineno; return FUNC;}
(return)/[^a-zA-Z0-9_]		  {yylloc.last_line=yylloc.first_line =yylineno; return RETURN;}
[\[\]\|\{\},\(\);\n]          {yylloc.last_line=yylloc.first_line =yylineno; return *yytext;}
[a-zA-Z_][a-zA-Z_0-9]*        {yylval.string=yytext; yylloc.last_line=yylloc.first_line =yylineno; return NAME;}

[ \t]				          {};
.				              {yylloc.last_line=yylloc.first_line =yylineno; yyerror(text);}

%%
