%{
#include "Nodes.h"
#include "lex.yy.c"

std::shared_ptr<Node> root;

int yylex(void);
void yyerror(const char*);
}

%union {
std::shared_ptr<Node> Node_;
int ival_;
short sval_;
std::string string_;
Logic bval_;
Datatypes types_;
std::vector<std::shared_ptr<Node>> opers_;
std::pair<Datatypes, std::string> param_;
std::vector<std::pair<Datatypes, std::string>> params_;
}

%token <ival_> INTVAL
%token <bval_> TRUE <bval_> FALSE <bval_> UNDEFINED
%token <types_> INT <types_> SHORT <types_> BOOL <types_> VECTOR
%token <string> NAME
%token DO WHILE IF
%nonassoc THEN
%nonassoc IFX
%nonassoc ELSE
%token BEGIN_ END
%left OR
%left AND
%left NOT
%token MOVE LEFT RIGHT
%token SET
%token LMS, OF
%token FUNC RETURN
%token SIZEOF
%left LARGER SMALLER
%left ADD SUB
%nonassoc '{' '}'

%type <> program
%type <> stmt_list
%type <> stmt
%type <> arglist
%type <> fargs
%type <> decl
%type <> varlist
%type <> varlistwithset
%type <> vecdecl
%type <> vecof
%type <> vecopr
%type <> expr
%type <> callargs
%type <> indexes
%type <> sizeofargs
%type <> type 
%type <> directions
%type <> operands
%type <> oplist


%%

program:
	'\n'												{}
	| program stmt '\n'									{}
	| program error '\n'								{}
	|													{}
	;


stmt_list:
	stmt												{}
	| stmt_list stmt									{}
	;

stmt:
	BEGIN_ END											{}
	| BEGIN_ stmt_list END								{}
	| ';'												{}
	| expr ';'											{}
	| decl ';'											{}
	| expr SET expr ';'									{}
	| DO stmt WHILE expr ';'							{}
	| IF expr THEN stmt %prec IFX						{}
	| IF expr THEN stmt ELSE stmt						{}
	| FUNC NAME argslist BEGIN_ stmt_list END			{}	
	;

argslist:
	'(' ')'												{}
	| '(' args ')'										{}
	;

fargs:
	type NAME											{}
	| args ',' type NAME								{}
	;

decl:
	varlist												{}
	| varlistwithset									{}
	| vecof type NAME SET '{' vecdecl '}' 				{}
	| vecof type NAME indexes							{}
	| vecof type NAME indexes SET '{' vecdecl '}'		{}
	;

varlist:
	type NAME											{}
	| varlist, NAME										{}
	;

varlistwithset:
	type NAME SET expr									{}
	| varlistwithset NAME SET expr						{}
	;

vecdecl:
	'{' expr '}'										{}
	| vecdecl ',' '{' vecdecl '}'						{}
	| vecopr											{}
	;

vecof:
	TYPE OF												{}
	| vecof TYPE OF										{}
	;

vecopr:
	expr												{}
	| vecopr ',' expr									{}
	;

expr:
	INT													{}
	| NAME												{}
	| TRUE												{}
	| FALSE												{}
	| UNDEFINED											{}
	| expr ADD expr										{}
	| expr SUB expr										{}
	| expr OR expr										{}
	| expr NOT OR expr									{}
	| expr AND expr										{}
	| expr NOT AND expr									{}
	| expr '|' expr SMALLER								{}
	| expr '|' expr LARGER								{}
	| '(' expr ')'										{}
	| SIZEOF '(' sizeofargs	')'							{}
	| NAME indexes										{}
	| directions										{}
	| NAME '(' callargs ')'								{}
	;

callargs:
	NAME												{}
	| callargs NAME										{}
	;

indexes:
	'[' expr ']'										{}
	| indexes '[' expr ']'								{}
	;

sizeofargs:
	type												{}
	| NAME												{}
	;

type:
	INT													{}
	| SHORT												{}
	| BOOL												{}
	| VECTOR											{}
	;

directions:
	MOVE RIGHT											{}
	| MOVE LEFT											{}
	| MOVE												{}
	| LEFT												{}
	| RIGHT												{}
	;

operands:
	'(' ')'												{}
	| '( oplist ')'										{}
	;

oplist:
	expr												{}
	| oplist expr										{}
	;

%%
