#include "Nodes.h"

typedef struct {
    int x;
    int y;
} coordinates;

std::vector<std::vector<std::string>> labirint; //0-тупик, 1-текущее, 2-пустота, 3-выходы
coordinates cur_coordinates;
int radius;

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

std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Node> u, std::shared_ptr<Memory> m) {
    switch (u->get_type())
    {
        case NodeType::VARLEAF: {
            try {
                return (*m)[std::dynamic_pointer_cast<VarLeaf>(u)->name];
            }
            catch (const std::exception&) {
                throw NameError("Variable with this name does not exist");
            }
        }
        case NodeType::INTLEAF: {
            try {
                return std::dynamic_pointer_cast<IntLeaf>(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::SHORTLEAF: {
            try {
                return std::dynamic_pointer_cast<ShortLeaf>(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::BOOLLEAF: {
            try {
                return std::dynamic_pointer_cast<BoolLeaf>(u)->data;
            }
            catch (std::exception& ex) {
                throw ex.what();
            }
        }
        case NodeType::ADDNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<AddNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<AddNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->add(*e2)) : throw SyntaxError("Add must have 2 parameters");
        }
        case NodeType::SUBNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<SubNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<SubNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->sub(*e2)) : throw SyntaxError("Sub must have 2 parameters");
        }
        case NodeType::ANDNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<AndNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<AndNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_and(*e2)) : throw SyntaxError("And must have 2 parameters");
        }
        case NodeType::NANDNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<NandNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<NandNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_nand(*e2)) : throw SyntaxError("Nand must have 2 parameters");
        }
        case NodeType::ORNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<OrNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<OrNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_or(*e2)) : throw SyntaxError("Or must have 2 parameters");
        }
        case NodeType::NORNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<NorNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<NorNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->logic_nor(*e2)) : throw SyntaxError("Nor must have 2 parameters");
        }
        case NodeType::SMALLERNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<SmallerNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<SmallerNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->smaller(*e2)) : throw SyntaxError("Smaller must have 2 parameters");
        }
        case NodeType::LARGERNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<LargerNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<LargerNode>(u)->right, m);
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->larger(*e2)) : throw SyntaxError("Larger must have 2 parameters");
        }
        case NodeType::SETNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<SetNode>(u)->left, m);
            auto e2 = exec(std::dynamic_pointer_cast<SetNode>(u)->right, m);
            if ((e1->data->get_type() == Datatypes::VECTOR && e2->data->get_type() != Datatypes::VECTOR) ||
            (e1->data->get_type() != Datatypes::VECTOR && e2->data->get_type() == Datatypes::VECTOR)) {
                throw TypeError("Types of lvalue and rvalue must be equal");
            }
            return (e1 && e2) ? std::make_shared<MemoryUnit>(e1->set(*e2)) : throw SyntaxError("Set must have lvalue and rvalue");
        }
        case NodeType::SIZEOFNODE: {
            auto e1 = exec(std::dynamic_pointer_cast<SizeofNode>(u)->next, m);
            return (e1) ? std::make_shared<MemoryUnit>(e1->size_of()) : throw SyntaxError("Sizeof must have variable or type");
        }
        case NodeType::IFNODE: {
            auto i = std::dynamic_pointer_cast<IfNode>(u);
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
                while ((bool)exec(std::dynamic_pointer_cast<LoopNode>(u)->condition, u->local)->data) {
                    exec(std::dynamic_pointer_cast<LoopNode>(u)->code, u->local);
                }
                u->local->clear();
                return nullptr;
            }
            catch (const CastError& c) {
                throw c;
            }
        }
        case NodeType::VARDECLNODE: {
            auto v = std::dynamic_pointer_cast<VarDeclNode>(u);
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
            auto v = std::dynamic_pointer_cast<VarListNode>(u);
            try {
                for (auto it : v->vec) {
                    exec(it, m);
                }
            } catch (std::exception ex) {
                throw ex;
            }
        }
        case NodeType::FDECLNODE: {
            auto f = std::dynamic_pointer_cast<FDeclNode>(u);
            m->add(f->func);
            u->local = m;
            return nullptr;
        }
        case NodeType::VECDECLNODE: {
            u->local =  m;
            auto vecn = std::dynamic_pointer_cast<VecDeclNode>(u);
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
                auto f = std::dynamic_pointer_cast<FcallNode>(u);
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
            auto i = std::dynamic_pointer_cast<IndexNode>(u);
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
            for (auto v : std::dynamic_pointer_cast<StatementList   >(u)->vec)
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
            } else {
                return std::make_shared<MemoryUnit>(std::make_shared<Bool>(UNDEFINED));
            }
            if (left < right && right <= radius) {
                if (right <= INT16_MAX) {
                    return std::make_shared<MemoryUnit>(std::make_shared<Short>(right));
                } else {
                    return std::make_shared<MemoryUnit>(std::make_shared<Int>(right));
                }
            } else {
                return std::make_shared<MemoryUnit>(std::make_shared<Bool>(UNDEFINED));
            }
        }
        default:
            break;
    }
}

