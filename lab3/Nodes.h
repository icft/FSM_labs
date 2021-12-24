#pragma once
#ifndef INTERPRETER_NODES_H
#define INTERPRETER_NODES_H

#include <utility>

#include "Memory.h"

class Node {
public:
    Node() = default;
    virtual ~Node() = default;
    virtual std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) = 0;
};

class Leaf {
public:
    std::shared_ptr<MemoryUnit> data;
    Leaf() = default;
    virtual ~Leaf() = default;
};

class IntLeaf : public Leaf {
public:
    IntLeaf() = default;
    explicit IntLeaf(const MemoryUnit& m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    explicit IntLeaf(int value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Int>(value));
    }
    ~IntLeaf() override = default;
};

class ShortLeaf : public Leaf {
public:
    ShortLeaf() = default;
    explicit ShortLeaf(const MemoryUnit& m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    explicit ShortLeaf(short value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Short>(value));
    }
    ~ShortLeaf() override = default;
};

class BoolLeaf : public Leaf {
public:
    BoolLeaf() = default;
    explicit BoolLeaf(const MemoryUnit& m) {
        data = std::make_shared<MemoryUnit>(m);
    }
    explicit BoolLeaf(Logic value) {
        data = std::make_shared<MemoryUnit>(std::make_shared<Bool>(value));
    }
    ~BoolLeaf() override = default;
};

class Variable : public Leaf {
private:
    std::string name;
public:
    Variable() = default;
    explicit Variable(std::string s) : name(std::move(s)) {}
    std::shared_ptr<MemoryUnit> getUnit(const std::shared_ptr<Memory>& m) {
        return m->getVariableUnit(name);
    }
    ~Variable() override = default;
};

class AddNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    AddNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->add(*u2));
        } else {
            //throw
        }
    }
    ~AddNode() override = default;
};

class SubNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    SubNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->sub(*u2));
        } else {
            //throw
        }
    }
    ~SubNode() override = default;
};

class AndNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    AndNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->logic_and(*u2));
        } else {
            //throw
        }
    }
    ~AndNode() override = default;
};

class NandNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    NandNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->logic_nand(*u2));
        } else {
            //throw
        }
    }
    ~NandNode() override = default;
};

class OrNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    OrNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->logic_or(*u2));
        } else {
            //throw
        }
    }
    ~OrNode() override = default;
};

class NorNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    NorNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->logic_nor(*u2));
        } else {
            //throw
        }
    }
    ~NorNode() override = default;
};

class SmallerNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    SmallerNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->smaller(*u2));
        } else {
            //throw
        }
    }
    ~SmallerNode() override = default;
};

class LargerNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    LargerNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        } else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->larger(*u2));
        } else {
            //throw
        }
    }
    ~LargerNode() override = default;
};

class SetNode : public Node {
private:
    std::shared_ptr<Node> left;
    std::shared_ptr<Node> right;
public:
    SetNode(const std::shared_ptr<Node>& l, const std::shared_ptr<Node>& r) {
        if (l && r) {
            left = l;
            right = r;
        }
        else {
            //throw
        }
    }
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> m) override {
        auto u1 = left->exec(m);
        auto u2 = right->exec(m);
        if (u1 && u2) {
            return std::make_shared<MemoryUnit>(u1->larger(*u2));
        }
        else {
            //throw
        }
    }
    ~SetNode() override = default;
};


class IfNode : public Node {
private:
    std::shared_ptr<Node> cond;
    std::shared_ptr<Node> if_code;
    std::shared_ptr<Node> else_code;
public:
    IfNode(std::shared_ptr<Node> c, std::shared_ptr<Node> i, std::shared_ptr<Node> e = nullptr) : 
        cond(c), if_code(i), else_code(e) {}
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> mem) {
        //вставить ошибку
        auto con = cond->exec(mem);
        auto internalMemory = std::make_shared<Memory>(mem);
        if (con) {
            if (con->data->to_bool() == Logic::TRUE) {
                if_code->exec(internalMemory);
            } else {
                else_code->exec(internalMemory);
            }
            return nullptr;
        }
    }
    ~IfNode() = default;
};

class Loop : public Node {
private:
    std::shared_ptr<Node> cond;
    std::shared_ptr<Node> code;
public:
    Loop(std::shared_ptr<Node> c, std::shared_ptr<Node> i) : cond(c), code(i) {}
    std::shared_ptr<MemoryUnit> exec(std::shared_ptr<Memory> mem) {
        //вставить ошибку
        auto internalMemory = std::make_shared<Memory>(mem);
        while (cond->exec(mem)->data->to_bool() == Logic::TRUE) {
            code->exec(internalMemory);
        }
        return nullptr;
    }
    ~Loop() = default;
};

class FuncNode : public Node {
private:
    std::string name;
    std::shared_ptr<Node> code;
    std::shared_ptr<FunctionUnit> f;

};

#endif //INTERPRETER_NODES_H
