#pragma once

#include "Memory.h"
#include <memory>
#include <map>
#include <vector>
#include <fstream>
#include <string>

std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Node> u, std::shared_ptr<Memory> m);

enum NodeType {VARLEAF, INTLEAF, SHORTLEAF, BOOLLEAF, ADDNODE, SUBNODE, ANDNODE, NANDNODE,
        ORNODE, NORNODE, SMALLERNODE, LARGERNODE, SETNODE, LOOPNODE, FDECLNODE,
        SIZEOFNODE, IFNODE, FCALLNODE, VECDECLNODE, INDEXNODE, VARDECLNODE, STATEMENT, VARLIST,
        RIGHTNODE, LEFTNODE, MOVENODE, LMSNODE};

class Node {
public:
    int line_number;
    std::shared_ptr<Memory> local;

    Node() = default;
    virtual ~Node() = default;
    virtual NodeType get_type() {}
};

//std::map<std::shared_ptr<Node>, std::shared_ptr<Memory>> stack;


class VarLeaf : public Node {
public:
    std::string name;
    NodeType type = NodeType::INTLEAF;

    VarLeaf() = default;
    VarLeaf(int line, const std::string& s) {
        line_number = line;
        name = s;
    }
    virtual ~VarLeaf() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class IntLeaf : public Node {
public:
    std::shared_ptr<MemoryUnit> data;
    NodeType type = NodeType::INTLEAF;

    IntLeaf() = default;
    IntLeaf(MemoryUnit m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    IntLeaf(int line, std::int32_t value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Int>(value));
        line_number = line;
    }
    virtual ~IntLeaf() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class ShortLeaf : public Node {
public:
    std::shared_ptr<MemoryUnit> data;
    NodeType type = NodeType::SHORTLEAF;

    ShortLeaf() = default;
    ShortLeaf(MemoryUnit m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    ShortLeaf(int line, std::int16_t value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Short>(value));
        line_number = line;
    }
    virtual ~ShortLeaf() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class BoolLeaf : public Node {
public:
    std::shared_ptr<MemoryUnit> data;
    NodeType type = NodeType::BOOLLEAF;

    BoolLeaf() = default;
    BoolLeaf(MemoryUnit m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    BoolLeaf(int line, Logic value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Bool>(value));
        line_number = line;
    }
    virtual ~BoolLeaf() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class AddNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::ADDNODE;

    AddNode() = default;
    AddNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Add must have 2 parameters");
        }
    }
    virtual ~AddNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class SubNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::SUBNODE;

    SubNode() = default;
    SubNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Sub must have 2 parameters");
        }
    }
    virtual ~SubNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class AndNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::ANDNODE;

    AndNode() = default;
    AndNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("And must have 2 parameters");
        }
    }
    virtual ~AndNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class NandNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::NANDNODE;

    NandNode() = default;
    NandNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Nand must have 2 parameters");
        }
    }
    virtual ~NandNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class OrNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::ORNODE;

    OrNode() = default;
    OrNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Or must have 2 parameters");
        }
    }
    virtual ~OrNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class NorNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::NORNODE;

    NorNode() = default;
    NorNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Nor must have 2 parameters");
        }
    }
    virtual ~NorNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class SmallerNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::SMALLERNODE;

    SmallerNode() = default;
    SmallerNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Smaller must have 2 parameters");
        }
    }
    virtual ~SmallerNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class LargerNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::LARGERNODE;

    LargerNode() = default;
    LargerNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Larger must have 2 parameters");
        }
    }
    virtual ~LargerNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class SetNode : public Node {
public:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
    NodeType type = NodeType::SETNODE;

    SetNode() = default;
    SetNode(int line, const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            line_number = line;
            left = l;
            right = r;
        }
        else {
            throw SyntaxError("Set must have lvalue and rvalue");
        }
    }
    virtual ~SetNode() = default;
};

class LoopNode : public Node {
public:
    std::shared_ptr<Node> condition;
    std::shared_ptr<Node> code;
    NodeType type = NodeType::LOOPNODE;

    LoopNode() = default;
    LoopNode(int line, std::shared_ptr<Node> cond, std::shared_ptr<Node> c) {
        if (cond == nullptr) {
            throw SyntaxError("Loop must have  a condition");
        }
        else {
            line_number = line;
            condition = cond;
            code = c;
        }
    }
    ~LoopNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class FDeclNode : public Node {
public:
    std::string name;
    std::shared_ptr<Node> code;
    std::shared_ptr<FunctionUnit> func;
    std::vector<std::pair<Datatypes, std::string>> params;
    std::shared_ptr<Node> ret;
    NodeType type = NodeType::FDECLNODE;

    FDeclNode() = default;
    FDeclNode(int line, std::string n, std::shared_ptr<Node> c, std::vector<std::pair<Datatypes, std::string>> p, std::shared_ptr<Node> r) {
        line_number = line;
        name = n;
        code = c;
        params = p;
        ret = r;
        func = std::make_shared<FunctionUnit>(name, params);
        func->set_link(std::shared_ptr<FDeclNode>(this));
    }
    virtual ~FDeclNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class FcallNode : public Node {
public:
    std::string name;
    std::vector<std::shared_ptr<Node>> params;
    NodeType type = NodeType::FCALLNODE;

    FcallNode() = default;
    FcallNode(int line, std::string n, std::vector<std::shared_ptr<Node>> p) : name(n), params(p) {
        line_number = line;
    }
    virtual ~FcallNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class SizeofNode : public Node {
public:
    std::shared_ptr<Node> next;
    NodeType type = NodeType::SIZEOFNODE;

    SizeofNode() = default;
    SizeofNode(int line, Datatypes t) {
        if ( t == Datatypes::INT) {
            line_number = line;
            next = std::make_shared<IntLeaf>(MemoryUnit(std::make_shared<Int>()));
        } else if (t == Datatypes::SHORT) {
            line_number = line;
            next = std::make_shared<IntLeaf>(MemoryUnit(std::make_shared<Short>()));
        } else if (t ==Datatypes::BOOL) {
            line_number = line;
            next = std::make_shared<IntLeaf>(MemoryUnit(std::make_shared<Bool>()));
        } else {
            throw SyntaxError("Can't be computed for a vector");
        }
    }
    SizeofNode(int line, std::string n) {
        line_number = line;
        next = std::make_shared<VarLeaf>(line, n);
    }
    virtual ~SizeofNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class IfNode : public Node {
public:
    std::shared_ptr<Node> condition;
    std::shared_ptr<Node> if_code;
    std::shared_ptr<Node> else_code;
    NodeType type = NodeType::IFNODE;

    IfNode() = default;
    IfNode(int line, std::shared_ptr<Node> c, std::shared_ptr<Node> i, std::shared_ptr<Node> e = nullptr) {
        if (c) {
            line_number = line;
            condition = c;
            if_code = i;
            else_code = e;
        }
        else {
            throw SyntaxError("If must have a condition");
        }
    }
    virtual ~IfNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class VecDeclNode :public Node {
public:
    std::string name;
    std::vector<std::shared_ptr<Node>> elems;
    std::vector<std::shared_ptr<Node>> dims;
    int vecof_count;
    bool main;

    std::vector<std::shared_ptr<Object>> objects;
    std::vector<int> real_dims;

    NodeType type = NodeType::VECDECLNODE;

    VecDeclNode() = default;
    VecDeclNode(int line, int v, std::string n, std::vector<std::shared_ptr<Node>> e, std::vector<std::shared_ptr<Node>> d) : vecof_count(v), name(n), elems(e), dims(d), main(true) {
        line_number = line;
    }
    VecDeclNode(int line, std::vector<std::shared_ptr<Node>> e) : elems(e) {
        name = {}; dims = {}; vecof_count = {}; main = false;
        line_number = line;
    }
    void init(std::vector<std::shared_ptr<Node>> v, std::shared_ptr<Memory> m) {
        std::pair<int, std::vector<Object>> p;
        auto k = false;
        auto s = 0;
        for (auto it = v.begin(); it != v.end(); it++) {
            if ((*it)->get_type() == NodeType::VECDECLNODE) {
                if (it == v.begin()) {
                    s = std::dynamic_pointer_cast<VecDeclNode>(*it)->elems.size();
                } else {
                    if (s != std::dynamic_pointer_cast<VecDeclNode>(*it)->elems.size()) {
                        throw SyntaxError("Different sizes");
                    }
                }
                init(std::dynamic_pointer_cast<VecDeclNode>(*it)->elems, m);
                k = true;
            } else if (k && (*it)->get_type() != NodeType::VECDECLNODE) {
                throw SyntaxError("Vector error");
            } else if (it != v.begin() && !k && (*it)->get_type() != NodeType::VECDECLNODE) {
                throw SyntaxError("Vector error");
            } else {
                auto n = exec(*it, m);
                if (!n) {
                    throw SyntaxError("Error in initialization");
                }
                objects.push_back(n->data);
            }
        }
        if (k) {
            real_dims.push_back(s);
        }
    };
    virtual ~VecDeclNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class IndexNode : public Node {
public:
    std::shared_ptr<Node> next;
    std::vector<std::shared_ptr<Node>> elems;
    NodeType type = NodeType::INDEXNODE;

    IndexNode() = default;
    IndexNode(int line, std::shared_ptr<Node> n, std::vector<std::shared_ptr<Node>> e) : next(n), elems(e) {
        if (!next) {
            throw SyntaxError("Vector missed");
        }
        line_number = line;
    }
    virtual ~IndexNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

typedef struct {
    std::string name;
    std::shared_ptr<Node> init;
} VarDeclaration;

class VarDeclNode : public Node {
public:
    std::string name;
    std::shared_ptr<Node> init;
    std::shared_ptr<VariableUnit> var;
    NodeType type = NodeType::VARDECLNODE;

    VarDeclNode() = default;
    VarDeclNode(int line, std::string n, Datatypes t, std::shared_ptr<Node> init_node = nullptr) {
        line_number = line;
        std::shared_ptr<Object> obj;
        if (t == Datatypes::INT) {
            obj = std::make_shared<Int>();
        }
        else if (t == Datatypes::SHORT) {
            obj = std::make_shared<Short>();
        }
        else {
            obj = std::make_shared<Bool>();
        }
        var = std::make_shared<VariableUnit>(n, obj);
        init = init_node;
    }
    VarDeclNode(std::string n, std::shared_ptr<MemoryUnit> m) {
        name = n;
        var = std::make_shared<VariableUnit>(n, m);
    }
    virtual ~VarDeclNode() = default;
    virtual NodeType get_type() {
        return type;
    }
};

class VarListNode : public Node {
public:
    Datatypes t;
    std::vector<std::shared_ptr<Node>> vec;
    NodeType type = NodeType::VARLIST;

    VarListNode(int line, std::pair<Datatypes, std::vector<VarDeclaration>> p) {
        t = p.first;
        line_number = line;
        for (auto it : p.second) {
            auto e = std::make_shared<VarDeclNode>(line, it.name, t, it.init);
            vec.push_back(e);
        }
    }
    ~VarListNode() = default;
};

class StatementList : public Node {
public:
    std::vector<std::shared_ptr<Node>> vec;
    NodeType type = NodeType::STATEMENT;

    StatementList() = default;
    StatementList(std::shared_ptr<Node> n) {
        if (n)
            vec.push_back(n);
    }
    void add(std::shared_ptr<Node> n) {
        if (n)
            vec.push_back(n);
    }
    /*
    void init() {
        local = std::shared_ptr<Memory>();
    }*/
    virtual ~StatementList() = default;
};

typedef struct {
    int x;
    int y;
} coordinates;

std::vector<std::vector<std::string>> labirint; //0-тупик, 1-текущее, 2-пустота, 3-выходы
coordinates cur_coordinates;
int radius;

class RightNode : public Node {
public:
    NodeType type = NodeType::RIGHTNODE;

    RightNode() = default;
    ~RightNode() = default;
};

class LeftNode : public Node {
public:
    NodeType type = NodeType::LEFTNODE;

    LeftNode() = default;
    ~LeftNode() = default;
};

class MoveNode : public Node {
public:
    NodeType type = NodeType::MOVENODE;

    MoveNode() = default;
    ~MoveNode() = default;
};

class LmsNode : public Node {
public:
    NodeType type = NodeType::STATEMENT;

    LmsNode() = default;
    ~LmsNode() = default;
};

void init_labirint(std::string filename) {
    std::ifstream file("lab");
    std::string s;
    std::string prob = " ";
    std::vector<std::vector<char>> vec;
    int radius;
    int x = -1, y = -1;
    while(getline(file, s)) {
        y++;
        x = 0;
        if (s.size() == 1) {
            radius = std::atoi(s.c_str());
        }
        std::vector<char> v;
        for (char c : s) {
            if (c != ' ') {
                if (c == '1') {
                    cur_coordinates.x = x;
                    cur_coordinates.y =y;
                }
                v.push_back(c);
                x++;
            }
        }
        vec.push_back(v);
    }
}

void init_memory(std::shared_ptr<Node> n) {
    n->local = std::shared_ptr<Memory>();
}

std::shared_ptr<VarLeaf> cast_var(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<VarLeaf>(n);
}

std::shared_ptr<IntLeaf> cast_int(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<IntLeaf>(n);
}

std::shared_ptr<ShortLeaf> cast_short(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<ShortLeaf>(n);
}

std::shared_ptr<BoolLeaf> cast_bool(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<BoolLeaf>(n);
}

std::shared_ptr<AddNode> cast_add(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<AddNode>(n);
}

std::shared_ptr<SubNode> cast_sub(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<SubNode>(n);
}

std::shared_ptr<AndNode> cast_and(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<AndNode>(n);
}

std::shared_ptr<OrNode> cast_or(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<OrNode>(n);
}

std::shared_ptr<NandNode> cast_nand(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<NandNode>(n);
}

std::shared_ptr<NorNode> cast_nor(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<NorNode>(n);
}

std::shared_ptr<SmallerNode> cast_smaller(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<SmallerNode>(n);
}

std::shared_ptr<LargerNode> cast_larger(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<LargerNode>(n);
}

std::shared_ptr<SetNode> cast_set(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<SetNode>(n);
}

std::shared_ptr<SizeofNode> cast_sizeof(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<SizeofNode>(n);
}

std::shared_ptr<FDeclNode> cast_fdecl(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<FDeclNode>(n);
}

std::shared_ptr<FcallNode> cast_fcall(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<FcallNode>(n);
}

std::shared_ptr<LoopNode> cast_loop(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<LoopNode>(n);
}

std::shared_ptr<VarDeclNode> cast_vardecl(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<VarDeclNode>(n);
}

std::shared_ptr<VarListNode> cast_varlist(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<VarListNode>(n);
}

std::shared_ptr<VecDeclNode> cast_vecdecl(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<VecDeclNode>(n);
}

std::shared_ptr<IndexNode> cast_index(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<IndexNode>(n);
}

std::shared_ptr<StatementList> cast_stmt(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<StatementList>(n);
}

std::shared_ptr<IfNode> cast_if(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<IfNode>(n);
}

std::shared_ptr<RightNode> cast_right(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<RightNode>(n);
}

std::shared_ptr<LeftNode> cast_left(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<LeftNode>(n);
}

std::shared_ptr<MoveNode> cast_move(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<MoveNode>(n);
}

std::shared_ptr<LmsNode> cast_lms(std::shared_ptr<Node> n) {
    return std::dynamic_pointer_cast<LmsNode>(n);
}

std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Node> u, std::shared_ptr<Memory> m) {
    switch (u->get_type())
    {
        case NodeType::VARLEAF: {
            try {
                return (*m)[cast_var(u)->name];
            }
            catch (const std::exception&) {
                throw NameError("Variable with this name does not exist");
            }
        }
        case NodeType::INTLEAF: {
            try {
                return cast_int(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::SHORTLEAF: {
            try {
                return cast_short(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::BOOLLEAF: {
            try {
                return cast_bool(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::ADDNODE: {
            auto e1 = exec(cast_add(u)->left, m);
            auto e2 = exec(cast_add(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->add(*e2)) : throw SyntaxError("Add must have 2 parameters");
        }
        case NodeType::SUBNODE: {
            auto e1 = exec(cast_sub(u)->left, m);
            auto e2 = exec(cast_sub(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->sub(*e2)) : throw SyntaxError("Sub must have 2 parameters");
        }
        case NodeType::ANDNODE: {
            auto e1 = exec(cast_and(u)->left, m);
            auto e2 = exec(cast_and(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_and(*e2)) : throw SyntaxError("And must have 2 parameters");
        }
        case NodeType::NANDNODE: {
            auto e1 = exec(cast_nand(u)->left, m);
            auto e2 = exec(cast_nand(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_nand(*e2)) : throw SyntaxError("Nand must have 2 parameters");
        }
        case NodeType::ORNODE: {
            auto e1 = exec(cast_or(u)->left, m);
            auto e2 = exec(cast_or(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_or(*e2)) : throw SyntaxError("Or must have 2 parameters");
        }
        case NodeType::NORNODE: {
            auto e1 = exec(cast_nor(u)->left, m);
            auto e2 = exec(cast_nor(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_nor(*e2)) : throw SyntaxError("Nor must have 2 parameters");
        }
        case NodeType::SMALLERNODE: {
            auto e1 = exec(cast_smaller(u)->left, m);
            auto e2 = exec(cast_smaller(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->smaller(*e2)) : throw SyntaxError("Smaller must have 2 parameters");
        }
        case NodeType::LARGERNODE: {
            auto e1 = exec(cast_larger(u)->left, m);
            auto e2 = exec(cast_larger(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->larger(*e2)) : throw SyntaxError("Larger must have 2 parameters");
        }
        case NodeType::SETNODE: {
            auto e1 = exec(cast_set(u)->left, m);
            auto e2 = exec(cast_set(u)->right, m);
            if ((e1->data->get_type() == Datatypes::VECTOR && e2->data->get_type() != Datatypes::VECTOR) ||
            (e1->data->get_type() != Datatypes::VECTOR && e2->data->get_type() == Datatypes::VECTOR)) {
                throw TypeError("Types of lvalue and rvalue must be equal");
            }
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->set(*e2)) : throw SyntaxError("Set must have lvalue and rvalue");
        }
        case NodeType::SIZEOFNODE: {
            auto e1 = exec(cast_sizeof(u)->next, m);
            return (e1) ? std::make_shared<MemoryUnit>(e1->size_of()) : throw SyntaxError("Sizeof must have variable or type");
        }
        case NodeType::IFNODE: {
            auto i = cast_if(u);
            u->local = std::make_shared<Memory>(m);
            auto e = exec(i->condition, m);
            if (e) {
                if ((bool)e->data) {
                    exec(i->if_code, u->local);
                }
                else if (i->else_code) {
                    exec(i->else_code, u->local);
                }
                u->local->clear();
                return nullptr;
            }
            else {
                throw SyntaxError("Invalid if condition");
            }
        }
        case NodeType::LOOPNODE: {
            u->local = std::make_shared<Memory>(m);
            try {
                while ((bool)exec(cast_loop(u)->condition, u->local)->data) {
                    exec(cast_loop(u)->code, u->local);
                }
                u->local->clear();
                return nullptr;
            }
            catch (const CastError& c) {
                throw c;
            }
        }
        case NodeType::VARDECLNODE: {
            auto v = cast_vardecl(u);
            try {
                m->add(v->var);
                if (v->init) {
                    auto i = exec(v->init, m);
                    if(i) {
                        v->var = std::make_shared<VariableUnit>(v->name, i);
                    }
                    else {
                        throw SyntaxError("Declaration must have rvalue");
                    }
                }
            } catch (std::exception ex) {
                throw ex;
            }
        }
        case NodeType::VARLIST: {
            auto v = cast_varlist(u);
            try {
                for (auto it : v->vec) {
                    exec(it, m);
                }
            } catch (std::exception ex) {
                throw ex;
            }
        }
        case NodeType::FDECLNODE: {
            auto f = cast_fdecl(u);
            m->add(f->func);
            u->local = m;
            return nullptr;
        }
        case NodeType::VECDECLNODE: {
            u->local =  m;
            auto vecn = cast_vecdecl(u);
            std::vector<int> tmp;
            auto elems = vecn->elems;
            auto dims = vecn->dims;
            try {
                if (!elems.empty() && !dims.empty()) {
                    if (vecn->vecof_count != dims.size()) {
                        throw SyntaxError("The number of VECTOR OF must be equivalent to nesting the vector");
                    }
                    vecn->init(elems, m);
                    std::vector<int> tmp;
                    for (auto d : dims) {
                        auto s = exec(d, m);
                        if (!s) {
                            throw SyntaxError("Error in initialization");
                        }
                        tmp.push_back((int)*(s->data));
                    }
                    if (vecn->real_dims != tmp) {
                        throw SyntaxError("Error in initialization");
                    }
                    auto vec = std::make_shared<Vector>(tmp, vecn->objects);
                    auto var = std::make_shared<VariableUnit>(vecn->name, vec);

                } else if (elems.empty() && !dims.empty()) {
                    std::vector<int> tmp;
                    for (auto d : dims) {
                        auto s = exec(d, m);
                        if (!s) {
                            throw SyntaxError("Error in initialization");
                        }
                        tmp.push_back((int)*(s->data));
                    }
                    auto vec = std::make_shared<Vector>(tmp);
                    auto var = std::make_shared<VariableUnit>(vecn->name, vec);
                } else if (!elems.empty() && dims.empty()) {
                    if (vecn->vecof_count != dims.size()) {
                        throw SyntaxError("The number of VECTOR OF must be equivalent to nesting the vector");
                    }
                    vecn->init(elems, m);
                    auto vec = std::make_shared<Vector>(vecn->objects);
                    auto var = std::make_shared<VariableUnit>(vecn->name, vec);
                } else {
                    throw SyntaxError("Either a size field or a field for setting values is required");
                }
            } catch (std::exception ex) {
                throw ex;
            }
        }
        case NodeType::FCALLNODE: {
            try {
                std::vector<std::shared_ptr<MemoryUnit>> tmp;
                auto f = cast_fcall(u);
                for (auto p : f->params) {
                    if (p) {
                        auto q = exec(p, m);
                        if (q)
                            tmp.push_back(q);
                        else
                            throw TypeError("It's not a parametr");
                    }
                    else
                        throw TypeError("It's not a parametr");
                }
                auto fnode = std::dynamic_pointer_cast<FunctionUnit>((*m)[f->name]);
                if (fnode->params.size() != f->params.size()) {
                    throw SyntaxError("Parameters more or less than required");
                }
                std::shared_ptr<Memory> local = std::make_shared<Memory>(m);
                for (int i = 0; i < fnode->params.size(); i++) {
                    if (fnode->params[i].first != tmp[i]->data->get_type()) {
                        throw TypeError("Type mismatch");
                    }
                    std::pair<std::string, std::shared_ptr<MemoryUnit>> p;
                    auto alloc = std::make_shared<VariableUnit>(fnode->params[i].second, tmp[i]);
                    local->add(alloc);
                }
                auto fdecl = std::dynamic_pointer_cast<FDeclNode>(fnode->link);
                std::shared_ptr<MemoryUnit> tmp1;
                if (fdecl->code) {
                    exec(fdecl->code, local);
                }
                if (fdecl->ret) {
                    tmp1 = exec(fdecl->ret, local);
                    local->clear();
                    return tmp1;
                } else {
                    throw SyntaxError("Return mismatch");
                }
            } catch (std::exception ex) {
                throw ex;
            }
        }
        case NodeType::INDEXNODE: {
            auto i = cast_index(u);
            auto v = exec(i->next, m);
            if (!v) {
                throw SyntaxError("Incorrect vector");
            }
            std::vector<int> tmp;
            for (auto p : i->elems) {
                auto t = exec(p, m);
                if (t)
                    tmp.push_back((int)*(t->data));
                else
                    throw SyntaxError("Incorrect path in vector");
            }
            return std::make_shared<MemoryUnit>((*v)[tmp]);
        }
        case NodeType::STATEMENT: {
            for (auto v : cast_stmt(u)->vec)
                exec(v, m);
            return nullptr;
        }
        case NodeType::RIGHTNODE: {
            int new_x;
            int new_y;
            if (cur_coordinates.y % 2 == 0 && cur_coordinates.x != labirint[0].size()-1) {
                new_x = cur_coordinates.x + 1;
                new_y = cur_coordinates.y + 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(1));
                }
            } else if (cur_coordinates.y % 2 != 0 && cur_coordinates.x != labirint[0].size()-1) {
                new_x = cur_coordinates.x;
                new_y = cur_coordinates.y - 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(1));
                }
            }
            return std::make_shared<MemoryUnit>(std::make_shared<Bool>(UNDEFINED));
        }
        case NodeType::LEFTNODE: {
            int new_x;
            int new_y;
            if (cur_coordinates.y % 2 == 0 && cur_coordinates.x != 0) {
                new_x = cur_coordinates.x + 1;
                new_y = cur_coordinates.y + 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(-1));
                }
            } else if (cur_coordinates.y % 2 != 0 && cur_coordinates.x != 0) {
                new_x = cur_coordinates.x + 1;
                new_y = cur_coordinates.y - 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(-1));
                }
            }
            return std::make_shared<MemoryUnit>(std::make_shared<Bool>(UNDEFINED));
        }
        case NodeType::MOVENODE: {
            int new_x;
            int new_y;
            if (cur_coordinates.y % 2 == 0 && cur_coordinates.y != 0) {
                new_x = cur_coordinates.x;
                new_y = cur_coordinates.y - 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(1));
                }
            } else if (cur_coordinates.y % 2 != 0 && cur_coordinates.x != labirint.size()-1) {
                new_x = cur_coordinates.x ;
                new_y = cur_coordinates.y + 1;
                if (labirint[new_y][new_x] != "0") {
                    cur_coordinates.x = new_x;
                    cur_coordinates.y = new_y;
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(-1));
                }
            }
            return std::make_shared<MemoryUnit>(std::make_shared<Bool>(UNDEFINED));
        }
        case NodeType::LMSNODE : {
            coordinates tmp_coordinates;
            int left = 0, right = 0;
            tmp_coordinates.x = cur_coordinates.x; tmp_coordinates.y = cur_coordinates.y;
            int new_x = tmp_coordinates.y,  new_y = tmp_coordinates.x;
            bool finish_left = false, finish_right = false;;
            while (labirint[new_y][new_x] != "0") {
                if (tmp_coordinates.y % 2 == 0 && tmp_coordinates.x != 0) {
                    new_x = tmp_coordinates.x + 1;
                    new_y = tmp_coordinates.y + 1;
                    if (labirint[new_y][new_x] == "3") {
                        finish_left = true;
                        break;
                    }
                    if (labirint[new_y][new_x] != "0") {
                        tmp_coordinates.x = new_x;
                        tmp_coordinates.y = new_y;
                        left++;
                    }
                } else if (tmp_coordinates.y % 2 != 0 && tmp_coordinates.x != 0) {
                    new_x = tmp_coordinates.x + 1;
                    new_y = tmp_coordinates.y - 1;
                    if (labirint[new_y][new_x] == "3") {
                        finish_left = true;
                        break;
                    }
                    if (labirint[new_y][new_x] != "0") {
                        tmp_coordinates.x = new_x;
                        tmp_coordinates.y = new_y;
                        left++;
                    }
                }
            }
            tmp_coordinates.x = cur_coordinates.x; tmp_coordinates.y = cur_coordinates.y;
            new_x = tmp_coordinates.y, new_y = tmp_coordinates.x;
            while (labirint[new_y][new_x] != "0") {
                if (tmp_coordinates.y % 2 == 0 && tmp_coordinates.x != 0) {
                    new_x = tmp_coordinates.x + 1;
                    new_y = tmp_coordinates.y + 1;
                    if (labirint[new_y][new_x] == "3") {
                        finish_right = true;
                        break;
                    }
                    if (labirint[new_y][new_x] != "0") {
                        tmp_coordinates.x = new_x;
                        tmp_coordinates.y = new_y;
                        right++;
                    }
                } else if (tmp_coordinates.y % 2 != 0 && tmp_coordinates.x != 0) {
                    new_x = tmp_coordinates.x + 1;
                    new_y = tmp_coordinates.y - 1;
                    if (labirint[new_y][new_x] == "3") {
                        finish_right = true;
                        break;
                    }
                    if (labirint[new_y][new_x] != "0") {
                        tmp_coordinates.x = new_x;
                        tmp_coordinates.y = new_y;
                        right++;
                    }
                }
            }
            if (finish_left) {
                if (right <= INT16_MAX) {
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(right));
                } else {
                    return std::make_shared<MemoryUnit>(std::make_shared<Int>(right));
                }
            }
            if (finish_right) {
                if (left <= INT16_MAX) {
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(left));
                } else {
                    return std::make_shared<MemoryUnit>(std::make_shared<Int>(left));
                }
            }
            if (left > right && left <= radius) {
                if (left <= INT16_MAX) {
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(-1*left));
                } else {
                    return std::make_shared<MemoryUnit>(std::make_shared<Int>(-1*left));
                }
            }
            if (left < right && right <= radius) {
                if (right <= INT16_MAX) {
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(right));
                } else {
                    return std::make_shared<MemoryUnit>(std::make_shared<Int>(right));
                }
            }
        }
        default:
            break;
    }
}
