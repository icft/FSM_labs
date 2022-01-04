%{
#include "Datatypes.h"
#include "Nodes.h"
#include "lex.yy.c"

std::shared_ptr<Node>* root;
extern int yylineno;
int yylex();
void yyerror(const char*);
%}

%union {
int ival;
std::string* s_name;
Logic bval;
Datatypes types_;
std::shared_ptr<Node>* Node_;
std::vector<std::shared_ptr<Node>>* lst;
std::vector<std::pair<Datatypes, std::string>>* params;
std::pair<Datatypes, std::vector<VarDeclaration>>* vd_list;
}

%token WORK
%token <ival> INTVAL
%token <bval> TRUE_ <bval> FALSE_ <bval> UNDEFINED_
%token <types_> INT_ <types_> SHORT_ <types_> BOOL_ <types_> VECTOR_
%token <s_name> NAME
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
%nonassoc '{' '}' '[' ']' '(' ')'

%type <Node_> program
%type <Node_> stmt_list
%type <Node_> stmt
//%type <Node_> fcall
%type <params> argslist
%type <params> fargs
%type <Node_> decl
%type <vd_list> varlist
%type <ival> vecof
%type <lst> vecdecl
%type <lst> vecdecl_list
%type <lst> expr_list
%type <Node_> expr
%type <lst> callargs
%type <lst> indexes
%type <types_> type
%type <Node_> directions


%%
program:
    FUNC WORK '(' ')' stmt RETURN '1' ';'                   {init_memory(*$5); exec(*$5, (*$5)->local);}
	;

stmt_list:
	stmt												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<StatementList>(*$1);}
	| stmt_list stmt									    {cast_stmt(*$1)->add(*$2); $$=$1;}
	;

stmt:
	BEGIN_ END											    {$$=new std::shared_ptr<Node>();}
	| BEGIN_ stmt_list END								    {$$=new std::shared_ptr<Node>(); $$=$2;}
	| ';'												    {$$=new std::shared_ptr<Node>(); *$$=nullptr;}
	| expr ';'											    {$$=$1;}
	| decl ';'											    {$$=$1;}
    | directions ';'									    {$$=$1;}
	| expr SET expr ';'									    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SetNode>(yylineno, *$1, *$3);}
	| DO stmt WHILE expr ';'							    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LoopNode>(yylineno, *$4, *$2);}
	| IF expr THEN stmt %prec IFX						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(yylineno, *$2, *$4);}
	| IF expr THEN stmt ELSE stmt						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(yylineno, *$2, *$4, *$6);}
	| FUNC NAME argslist BEGIN_ stmt_list END RETURN expr ';' 	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<FDeclNode>(yylineno, *$2, *$5, *$3, *$8);}
	;

argslist:
	'(' ')'												    {$$=new std::vector<std::pair<Datatypes, std::string>>();}
	| '(' fargs ')'										    {$$=$2;}
	;

fargs:
	type NAME											    {$$=new std::vector<std::pair<Datatypes, std::string>>();
	                                                         std::pair<Datatypes, std::string> p; p.first=$1; p.second=*$2;
	                                                         $$->push_back(p);}
	| fargs ',' type NAME								    {std::pair<Datatypes, std::string> p; p.first=$3; p.second=*$4;
                                                             $1->push_back(p); $$=$1;}
	;

decl:
	varlist												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<VarListNode>(yylineno, *$1);}
	| vecof type NAME SET vecdecl       				    {$$=new std::shared_ptr<Node>(); std::vector<std::shared_ptr<Node>> v;
	                                                         *$$= std::make_shared<VecDeclNode>(yylineno, $1, *$3, *$5, v);}
	| vecof type NAME indexes							    {$$=new std::shared_ptr<Node>(); std::vector<std::shared_ptr<Node>> v;
	                                                         *$$= std::make_shared<VecDeclNode>(yylineno, $1, *$3, v, *$4);}
	| vecof type NAME indexes SET vecdecl          		    {$$=new std::shared_ptr<Node>(); *$$= std::make_shared<VecDeclNode>(yylineno, $1, *$3, *$6, *$4);}
	;

varlist:
	type NAME											    {$$=new std::pair<Datatypes, std::vector<VarDeclaration>>();
	                                                         (*$$).first = $1; VarDeclaration p; p.name = *$2; (*$$).second.push_back(p);}
    | type NAME SET expr									{$$=new std::pair<Datatypes, std::vector<VarDeclaration>>();
                                                             (*$$).first = $1; VarDeclaration p; p.name = *$2; p.init = *$4; (*$$).second.push_back(p);}
    | varlist ',' NAME										{VarDeclaration p; p.name = *$3; (*$1).second.push_back(p); $$=$1;}
	| varlist ',' NAME SET expr						        {VarDeclaration p; p.name = *$3; p.init = *$5; (*$1).second.push_back(p); $$=$1;}
    ;

vecof:
	VECTOR_ OF												{$$=1;}
	| vecof VECTOR_ OF										{$$=$1+1;}
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
	INTVAL													{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IntLeaf>(yylineno, $1);}
	| NAME												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<VarLeaf>(yylineno, *$1);}
	| TRUE_												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| FALSE_											    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| UNDEFINED_										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(yylineno, $1);}
	| expr ADD expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<AddNode>(yylineno, *$1, *$3);}
	| expr SUB expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SubNode>(yylineno, *$1, *$3);}
	| expr OR expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<OrNode>(yylineno, *$1, *$3);}
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
	| LMS                                                   {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LmsNode>();}
	;

callargs:
	expr											    	{$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1);}
	| callargs expr									    	{$1->push_back(*$2); $$=$1;}
	;

indexes:
	'[' expr ']'										    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$2);}
	| indexes '[' expr ']'								    {$1->push_back(*$3); $$=$1;}
	;

type:
	INT_												    {$$=$1;}
	| SHORT_											    {$$=$1;}
	| BOOL_												    {$$=$1;}
	| VECTOR_											    {$$=$1;}
	;

directions:
	MOVE RIGHT											    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<RightNode>();}
	| MOVE LEFT											    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LeftNode>();}
	| MOVE												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<MoveNode>();}
	| LEFT												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LeftNode>();}
	| RIGHT												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<RightNode>();}
	;
%%

void yyerror(const char* c) {
    std::string str;
    str = c;
    std::cout << str << std::endl;
}

int main() {
}

