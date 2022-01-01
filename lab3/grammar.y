%{
#include "Nodes.h"
#include "lex.yy.c"

std::shared_ptr<Node>* root;

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
%token <string_> NAME
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
%token LARGER SMALLER
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
    | directions ';'									    	{$$=$1;}
	| expr SET expr ';'									    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SetNode>(*$1, *$3);}
	| DO stmt WHILE expr ';'							    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LoopNode>(*$4, *$2);}
	| IF expr THEN stmt %prec IFX						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(*$2, *$4);}
	| IF expr THEN stmt ELSE stmt						    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IfNode>(*$2, *$4, *$6);}
	| FUNC NAME argslist BEGIN_ stmt_list END RETURN expr;	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<FDeclNode>(*$2, *$5, *$3, *$8);}
	;

argslist:
	'(' ')'												    {}
	| '(' fargs ')'										    {}
	;

fargs:
	type NAME											    {}
	| fargs ',' type NAME								    {}
	;

decl:
	varlist												    {}
	| vecof type NAME SET vecdecl       				    {}
	| vecof type NAME indexes							    {}
	| vecof type NAME indexes SET vecdecl          		    {}
	;

varlist:
	type NAME											    {}
    | type NAME SET expr									{}	
    | varlist ',' NAME										    {}
	| varlist ',' NAME SET expr						        {}
    ;

vecof:
	VECTOR OF												    {}
	| vecof VECTOR OF										    {}
	;
	
vecdecl:
	'{' expr '}'									    {}
	| '{' vecdecl_list '}'						    {}
	;

vecdecl_list:
	vecdecl 											    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1); delete $1;}
	| vecdecl_list ',' vecdecl									    {$1->push_back(*$3); $$=$1; delete $3;}
	;
	
expr:
	INT													    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IntLeaf>(*$1);}
	| SHORT                                                 {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<ShortLeaf>(*$1);}
	| NAME												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<VarLeaf>(*$1);}
	| TRUE												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(*$1);}
	| FALSE												    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(*$1);}
	| UNDEFINED											    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<BoolLeaf>(*$1);}
	| expr ADD expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<AddNode>(*$1, *$3);}
	| expr SUB expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SubNode>(*$1, *$3);}
	| expr OR expr										    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<OrNode>(*$1, *$3);}
	| expr NOT OR expr									    {$$=new std::shared_ptr<Node>(); *$$=std::make_shared<NorNode>(*$1, *$3);}
	| expr AND expr									    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<AndNode>(*$1, *$3);}
	| expr NOT AND expr								    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<NandNode>(*$1, *$3);}
	| expr '|' expr SMALLER							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SmallerNode>(*$1, *$3);}
	| expr '|' expr LARGER							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<LargerNode>(*$1, *$3);}
	| '(' expr ')'									    	{$$=$2;}
	| SIZEOF '(' sizeofargs	')'						    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<SizeofNode>(*$3);}
	| expr indexes									    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<IndexNode>(*$1, *$2);}
//	| directions									    	{$$=$1;}
	| NAME '(' callargs ')'							    	{$$=new std::shared_ptr<Node>(); *$$=std::make_shared<FcallNode>(*$1, *$3);}
	;

callargs:
	NAME											    	{$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1); delete $1;}
	| callargs NAME									    	{$1->push_back(*$2); $$=$1; delete $2;}
	;

indexes:
	'[' expr ']'										    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$2); delete $2;}
	| indexes '[' expr ']'								    {$1->push_back(*$3); $$=$1; delete $3;}
	;

sizeofargs:
	type											    	{$$=$1;}
	| NAME												    {$$=$1;}
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

operands:
	'(' ')'												    {$$=new std::vector<std::shared_ptr<Node>>();}
	| '(' oplist ')'										    {$$=$2;}
	;
oplist:
	expr												    {$$=new std::vector<std::shared_ptr<Node>>(); $$->push_back(*$1); delete $1;}
	| oplist expr										    {$1->push_back(*$2); $$=$1; delete $2;}
	;
%%
