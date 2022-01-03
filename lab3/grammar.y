%{
#include "Nodes.h"
#include "lex.yy.c"
#include <iostream>
#include <memory>
#include <utility>

std::shared_ptr<Node>* root;
extern int yylineno;
int yylex(void);
void yyerror(const char*);
%}

%union {
std::shared_ptr<Node>* Node_;
int ival;
short sval;
std::string string;
Logic bval;
Datatypes types_;
std::vector<std::shared_ptr<Node>>* lst;
std::pair<Datatypes, std::string> param;
std::vector<std::pair<Datatypes, std::string>>* params;
VarDeclaration vd;
std::pair<Datatypes, std::vector<VarDeclaration>>* vd_list;
std::shared_ptr<StatementList>* statement;
}

%token <ival> INTVAL
%token <bval> TRUE <bval> FALSE <bval> UNDEFINED
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
%token LMS OF
%token FUNC RETURN
%token SIZEOF
%token LARGER SMALLER
%left ADD SUB
%nonassoc '{' '}' '[' ']'

%type <statement> program
%type <statement> stmt_list
%type <Node_> stmt
%type <params> argslist
%type <params> fargs
%type <Node_> decl
%type <vd_list> varlist
%type <int> vecof
%type <lst> vecdecl
%type <lst> vecdecl_list
%type <lst> expr_list
%type <lst> expr
%type <lst> callargs
%type <lst> indexes
%type <types_> type
%type <Node_> directions
//%type <lst> operands
//%type <lst> oplist


%%

program:
	'\n'												    {}
	| program stmt                                          {(**$1).add(*$2); $$=$1; delete $2; root=$$;}
	| program stmt '\n'									    {(**$1).add(*$2); $$=$1; delete $2; root=$$;}
	| program error '\n'								    {}
	|													    {$$=new std::make_shared<Statement>(); (*$$)=std::make_shared<Statement>();root=$$;}
	;

stmt_list:
	stmt												    {$$=new std::Shared_ptr<Statement>(); *$$=std::make_shared<Statement>(*$1); delete $1;}
	| stmt_list stmt									    {$1->add($2); $$=$1; delete $2;}
	;

stmt:
	BEGIN_ END											    {$$=new std::shared_ptr<Node>();}
	| BEGIN_ stmt_list END								    {$$=new std::shared_ptr<Node>(); *$$=*$2; delete $2;}
	| ';'												    {$$=new std::shared_ptr<Node>(); *$$=nullptr;}
	| expr ';'											    {$$=$1;}
	| decl ';'											    {$$=$1;}
    | directions ';'									    {$$=$1;}
	| expr SET expr ';'									    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SetNode>(yylineno, *$1, *$3);}
	| DO stmt WHILE expr ';'							    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LoopNode>(yylineno, *$4, *$2);}
	| IF expr THEN stmt %prec IFX						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(yylineno, *$2, *$4);}
	| IF expr THEN stmt ELSE stmt						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(yylineno, *$2, *$4, *$6);}
	| FUNC NAME argslist BEGIN_ stmt_list END RETURN expr ';'	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<FDeclNode>(yylineno, *$2, *$5, *$3, *$8);}
	;

argslist:
	'(' ')'												    {$$=new std::vector<param_>();}
	| '(' fargs ')'										    {$$=$2;}
	;

fargs:
	type NAME											    {$$=new std::vector<param>();
	                                                         param p; p.first=*$1; p.second=*$2;
	                                                         $$->push_back(p);}
	| fargs ',' type NAME								    {param p; p.first=*$3; p.second=*$4;
                                                             $1->push_back(p); $$=$1;}
	;

decl:
	varlist												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<VarListNode>(yylineno, *$1);}
	| vecof type NAME SET vecdecl       				    {$$=new std::shared_ptr<VecDeclNode>(yylineno, $1, *$3, *$5, nullptr);}
	| vecof type NAME indexes							    {$$=new std::shared_ptr<VecDeclNode>(yylineno, $1, *$3, nullptr, *$4);}
	| vecof type NAME indexes SET vecdecl          		    {$$=new std::shared_ptr<VecDeclNode>(yylineno, $1, *$3, *$6, *$4);}
	;

varlist:
	type NAME											    {$$=new std::pair<Datatypes, std::vector<vd>>();
	                                                         (*$$).first = $1; vd p; p.name = *$2; (*$$).second->push_back(p);}
    | type NAME SET expr									{$$=new std::pair<Datatypes, std::vector<vd>>();
                                                             (*$$).first = $1; vd p; p.name = *$2; p.init = $4; (*$$).second->push_back(p);}
    | varlist ',' NAME										{vd p; p.name = *$3; $1->push_back(p); $$=$1;}
	| varlist ',' NAME SET expr						        {vd p; p.name = *$3; p.init = *$5; $1->push_back(p); $$=$1;}
    ;

vecof:
	VECTOR OF												{$$=1;}
	| vecof VECTOR OF										{$$=$1+1;}
	;
	
vecdecl:
	'{' expr_list '}'									    {$$=$2;}
	| '{' vecdecl_list '}'						            {$$=$2;}
	;

vecdecl_list:
	vecdecl 											    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(std::make_shared<VecDeclNode>(yylineno, *$1));}
	| vecdecl_list ',' vecdecl							    {$1->push_back(std::make_shared<VecDeclNode>(yylineno, *$3)); $$=$1;}
	;

expr_list:
    expr                                                    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1);}
    | expr_list ',' expr                                    {$1->push_back(*$3); $$=$1;}
    ;

expr:
	INT													    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IntLeaf>(yylineno, $1);}
	| SHORT                                                 {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<ShortLeaf>(yylineno,$1);}
	| NAME												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<VarLeaf>(yylineno, *$1);}
	| TRUE												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| FALSE												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| UNDEFINED											    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| expr ADD expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<AddNode>(yylineno, *$1, *$3);}
	| expr SUB expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SubNode>(yylineno, *$1, *$3);}
	| expr OR expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<OrNode>(*yylineno, $1, *$3);}
	| expr NOT OR expr									    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<NorNode>(yylineno, *$1, *$4);}
	| expr AND expr									    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<AndNode>(yylineno, *$1, *$3);}
	| expr NOT AND expr								    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<NandNode>(yylineno, *$1, *$4);}
	| expr '|' expr SMALLER							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SmallerNode>(yylineno, *$1, *$3);}
	| expr '|' expr LARGER							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LargerNode>(yylineno, *$1, *$3);}
	| '(' expr ')'									    	{$$=$2;}
	| SIZEOF '(' type ')'   						    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SizeofNode>(yylineno, $3);}
	| SIZEOF '(' NAME ')'                                   {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SizeofNode>(yylineno, *$3);}
	| expr indexes									    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IndexNode>(yylineno, *$1, *$2);}
	| NAME '(' callargs ')'							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<FcallNode>(yylineno, *$1, *$3);}
	;

callargs:
	NAME											    	{$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1); delete $1;}
	| callargs NAME									    	{$1->push_back(*$2); $$=$1; delete $2;}
	;

indexes:
	'[' expr ']'										    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$2); delete $2;}
	| indexes '[' expr ']'								    {$1->push_back(*$3); $$=$1; delete $3;}
	;

type:
	INT													    {$$=$1;}
	| SHORT												    {$$=$1;}
	| BOOL												    {$$=$1;}
	| VECTOR											    {$$=$1;}
	;

directions:
	MOVE RIGHT											    {}
	| MOVE LEFT											    {}
	| MOVE												    {}
	| LEFT												    {}
	| RIGHT												    {}
	;

//operands:
//	'(' ')'												    {$$=new //std::vector<std::shared_ptr<Node>>();}
//	| '(' oplist ')'										{$$=$2;}
//	;
//oplist:
//	expr												    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1); delete $1;}
//	| oplist expr										    {$1->push_back(*$2); $$=$1; delete $2;}
//	;
%%

void yyerrpr(const char* c) {
    std::string str;
    str = c;
    std::cout << str << std::endl;
}

void main() {
    fopen();

}
