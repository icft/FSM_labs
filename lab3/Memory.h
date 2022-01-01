#pragma once

#include "Datatypes.h"
#include <map>
#include <memory>
#include "Nodes.h"

enum class UnitType {VARIABLE, FUNCTION};

class MemoryUnit {
public:
    std::shared_ptr<Object> data;

    MemoryUnit() : data(nullptr) {}
    MemoryUnit(std::shared_ptr<Object> d, bool l = false) : data(d) {}
    MemoryUnit(MemoryUnit& m) : data(m.data) {}
    MemoryUnit(MemoryUnit&& m) : data(std::move(m.data)) {}
    MemoryUnit set(const MemoryUnit& m) {
        data = m.data;
    }
    MemoryUnit add(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->get_type() == Datatypes::SHORT && m.data->get_type() == Datatypes::SHORT) {
                auto s1 = std::dynamic_pointer_cast<Short>(data);
                auto s2 = std::dynamic_pointer_cast<Short>(m.data);
                if (s2->value >= 0) {
                    if (s1->value > INT16_MAX - s2->value) {
                        res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                    }
                    else {
                        res.data = std::make_shared<Short>((short)*data + (short)*m.data);
                    }
                }
                else {
                    if (s1->value < INT16_MIN - s2->value) {
                        res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                    }
                    else {
                        res.data = std::make_shared<Short>((short)*data + (short)*m.data);
                    }
                }
            }
            else if (data->get_type() == Datatypes::INT && m.data->get_type() == Datatypes::INT) {
                auto i1 = std::dynamic_pointer_cast<Int>(data);
                auto i2 = std::dynamic_pointer_cast<Int>(m.data);
                if (i2->value >= 0) {
                    if (i1->value > INT32_MAX - i2->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                    }
                }
                else {
                    if (i1->value < INT32_MIN - i2->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                    }
                }
            }
            else if (data->get_type() == Datatypes::SHORT && m.data->get_type() == Datatypes::INT) {
                auto s = std::dynamic_pointer_cast<Short>(data);
                auto i = std::dynamic_pointer_cast<Int>(m.data);
                if (i->value >= 0) {
                    if (s->value > INT32_MAX - i->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (s->value > INT16_MAX - i->value) {
                            res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data + (int)*m.data);
                        }
                    }
                }
                else {
                    if (s->value < INT32_MIN - i->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (s->value < INT16_MIN - i->value) {
                            res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data + (int)*m.data);
                        }
                    }
                }
            }
            else if (data->get_type() == Datatypes::INT && m.data->get_type() == Datatypes::SHORT) {
                auto i = std::dynamic_pointer_cast<Int>(data);
                auto s = std::dynamic_pointer_cast<Short>(m.data);
                if (s->value >= 0) {
                    if (i->value > INT32_MAX - s->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (i->value > INT16_MAX - s->value) {
                            res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data + (int)*m.data);
                        }
                    }
                }
                else {
                    if (i->value < INT32_MIN - s->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (i->value < INT16_MIN - s->value) {
                            res.data = std::make_shared<Int>((int)*data + (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data + (int)*m.data);
                        }
                    }
                }
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        //res.data = std::make_shared<Int>((int)(*data)+(int)(*m.data));
        return res;
    }
    MemoryUnit sub(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->get_type() == Datatypes::SHORT && m.data->get_type() == Datatypes::SHORT) {
                auto s1 = std::dynamic_pointer_cast<Short>(data);
                auto s2 = std::dynamic_pointer_cast<Short>(m.data);
                if (s2->value >= 0) {
                    if (s1->value < INT16_MIN + s2->value) {
                        res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                    }
                    else {
                        res.data = std::make_shared<Short>((short)*data - (short)*m.data);
                    }
                }
                else {
                    if (s1->value > INT16_MAX + s2->value) {
                        res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                    }
                    else {
                        res.data = std::make_shared<Short>((short)*data - (short)*m.data);
                    }
                }
            }
            else if (data->get_type() == Datatypes::INT && m.data->get_type() == Datatypes::INT) {
                auto i1 = std::dynamic_pointer_cast<Int>(data);
                auto i2 = std::dynamic_pointer_cast<Int>(m.data);
                if (i2->value >= 0) {
                    if (i1->value < INT32_MIN + i2->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                    }
                }
                else {
                    if (i1->value > INT32_MAX + i2->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                    }
                }
            }
            else if (data->get_type() == Datatypes::SHORT && m.data->get_type() == Datatypes::INT) {
                auto s = std::dynamic_pointer_cast<Short>(data);
                auto i = std::dynamic_pointer_cast<Int>(m.data);
                if (i->value >= 0) {
                    if (s->value < INT32_MIN + i->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (s->value < INT16_MIN + i->value) {
                            res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data - (int)*m.data);
                        }
                    }
                }
                else {
                    if (s->value > INT32_MAX + i->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (s->value > INT16_MAX + i->value) {
                            res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data - (int)*m.data);
                        }
                    }
                }
            }
            else if (data->get_type() == Datatypes::INT && m.data->get_type() == Datatypes::SHORT) {
                auto i = std::dynamic_pointer_cast<Int>(data);
                auto s = std::dynamic_pointer_cast<Short>(m.data);
                if (s->value >= 0) {
                    if (i->value < INT32_MIN + s->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (i->value > INT16_MIN + s->value) {
                            res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data - (int)*m.data);
                        }
                    }
                }
                else {
                    if (i->value > INT32_MAX + s->value) {
                        throw OverflowError("Int Overflow");
                    }
                    else {
                        if (i->value > INT16_MAX + s->value) {
                            res.data = std::make_shared<Int>((int)*data - (int)*m.data);
                        }
                        else {
                            res.data = std::make_shared<Short>((int)*data - (int)*m.data);
                        }
                    }
                }
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        //res.data = std::make_shared<Int>((int)(*data) - (int)(*m.data));
        return res;
    }
    MemoryUnit logic_and(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->to_bool() == Logic::UNDEFINED || m.data->to_bool() == Logic::UNDEFINED) {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
            else if (data->to_bool() == Logic::FALSE || m.data->to_bool() == Logic::FALSE) {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
        } catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit logic_or(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->to_bool() == Logic::UNDEFINED || m.data->to_bool() == Logic::UNDEFINED) {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
            else if (data->to_bool() == Logic::FALSE && m.data->to_bool() == Logic::FALSE) {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
        } catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit logic_nand(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->to_bool() == Logic::UNDEFINED || m.data->to_bool() == Logic::UNDEFINED) {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
            else if (data->to_bool() == Logic::TRUE && m.data->to_bool() == Logic::TRUE) {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit logic_nor(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if (data->to_bool() == Logic::UNDEFINED || m.data->to_bool() == Logic::UNDEFINED) {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
            else if (data->to_bool() == Logic::FALSE && m.data->to_bool() == Logic::FALSE) {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit smaller(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if ((int)(*data) > (int)(*m.data)) {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
            else if ((int)(*data) > (int)(*m.data)) {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit larger(const MemoryUnit& m) {
        MemoryUnit res;
        try {
            if ((int)(*data) > (int)(*m.data)) {
                res.data = std::make_shared<Bool>(Logic::TRUE);
            }
            else if ((int)(*data) < (int)(*m.data)) {
                res.data = std::make_shared<Bool>(Logic::FALSE);
            }
            else {
                res.data = std::make_shared<Bool>(Logic::UNDEFINED);
            }
        }
        catch (const std::exception& ex) {
            throw ex;
        }
        return res;
    }
    MemoryUnit size_of() {
        MemoryUnit res;
        if (data->get_type() == Datatypes::INT) {
            res.data = std::make_shared<Int>(32);
        }
        if (data->get_type() == Datatypes::SHORT) {
            res.data = std::make_shared<Int>(16);
        }
        if (data->get_type() == Datatypes::BOOL) {
            res.data = std::make_shared<Int>(16);
        }
        if (data->get_type() == Datatypes::VECTOR) {
            throw TypeError("Sizeof can't be used with type vector");
        }
        return res;
    }
    MemoryUnit operator[](std::vector<int> dims) {
        if (data->get_type() == Datatypes::VECTOR) {
            return MemoryUnit((*std::dynamic_pointer_cast<Vector>(data))[dims]);
        } else {
            throw TypeError("Index operator only for vector");
        }
    }
    virtual UnitType get_type() {}
    ~MemoryUnit() = default;
};

class VariableUnit : public MemoryUnit {
public:
    std::string name;
    UnitType type = UnitType::VARIABLE;

    ~VariableUnit() = default;
    VariableUnit(std::string s, std::shared_ptr<MemoryUnit> m) : name(s), MemoryUnit(*m) {}
    VariableUnit(std::string n, std::shared_ptr<Object> obj) : name(n), MemoryUnit(obj) {}
    std::string get_name() {
        return name;
    }
    virtual UnitType get_type() {
        return type;
    }
};

class Node;

class FunctionUnit : public MemoryUnit {
public:
    std::string name;
    std::vector<std::pair<Datatypes, std::string>> params;
    std::shared_ptr<Node> link;
    UnitType type = UnitType::FUNCTION;

    FunctionUnit(std::string n, std::vector<std::pair<Datatypes, std::string>> p) {
        name = n;
        params = p;
    }
    ~FunctionUnit() = default;
    void set_link(std::shared_ptr<Node> f) {
        link = f;
    }
    std::shared_ptr<Node> get_link() {
        return link;
    }
    std::string get_name() {
        return name;
    }
    bool check_type(Datatypes t, int number) {
        for (auto i = 0; i < params.size(); i++) {
            if (params[i].first == t && number == i + 1) {
                return true;
            }
            else {
                return false;
            }
        }
    }
    bool check_all(std::vector<Datatypes> t) {
        if (t.size() != params.size()) {
            return false;
        }
        for (auto i = 0; i < params.size(); i++) {
            if (params[i].first != t[i]) {
                return false;
            }
        }
        return true;
    }
    virtual UnitType get_type() {
        return type;
    }
};



class Memory {
private:
    std::shared_ptr<Memory> parent;
    std::map<std::string, std::shared_ptr<MemoryUnit>> local;
public:
    Memory(std::shared_ptr<Memory> p) : parent(p) {}
    Memory(const Memory& m) : parent(m.parent), local(m.local) {}
    void clear() {
        local.clear();
        parent = nullptr;
    }
    void add(std::shared_ptr<VariableUnit> var) {
        if (local.contains(var->get_name())) {
            throw MemoryError("Variable with this name exists");
        }
        local[var->get_name()] = var;
    }
    void add(std::shared_ptr<FunctionUnit> func) {
        if (local.contains(func->get_name())) {
            throw MemoryError("Function with this name exists");
        }
        local[func->get_name()] = func;
    }
    std::shared_ptr<VariableUnit> operator[](std::string name) {
        if (local[name]->get_type() != UnitType::VARIABLE) {
            throw MemoryError("Variable with this name doesn't exists");
        } else if (local.contains(name) && local[name]->get_type() == UnitType::VARIABLE) {
            return std::dynamic_pointer_cast<VariableUnit>(local[name]);
        } else {
            return (*parent)[name];
        }
    }
    std::shared_ptr<FunctionUnit> operator[](std::pair<std::string, std::vector<Datatypes>> p) {
        if (local[p.first]->get_type() != UnitType::FUNCTION) {
            throw MemoryError("Function with this name doesn't exists");
        } else if (local.contains(p.first) && local[p.first]->get_type() == UnitType::FUNCTION) {
            if (!std::dynamic_pointer_cast<FunctionUnit>(local[p.first])->check_all(p.second)) {
                throw TypeError("Error in arg list");
            } else {
                return std::dynamic_pointer_cast<FunctionUnit>(local[p.first]);
            }
        } else {
            return (*parent)[p];
        }
    }
};
