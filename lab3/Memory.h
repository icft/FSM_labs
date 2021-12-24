#pragma once
#ifndef INTERPRETER_MEMORY_H
#define INTERPRETER_MEMORY_H

#include <utility>

#include "DataTypes.h"
#include "Nodes.h"

class MemoryUnit {
public:
    std::shared_ptr<Object> data;
    bool locator;

    MemoryUnit() : data(nullptr), locator(false) {}
    MemoryUnit(std::shared_ptr<Object> d, bool l = false) : data(std::move(d)), locator(l) {}
    MemoryUnit(MemoryUnit& c) {
        data = c.data;
        locator = c.locator;
    }
    MemoryUnit(MemoryUnit&& c)  noexcept {
        data = std::move(c.data);
        locator = c.locator;
    }
    MemoryUnit set(const MemoryUnit& c) {
        data = c.data;
        locator = c.locator;
    }
    MemoryUnit add(const MemoryUnit& c) {
        MemoryUnit res;
        res.data = std::make_shared<Int>((int)(*data));
        return res;
    }
    MemoryUnit sub(const MemoryUnit& c) {
        MemoryUnit res;
        res.data = std::make_shared<Int>((int)(*data)-(int)(*c.data));
        return res;
    }
    MemoryUnit logic_and(const MemoryUnit& c) {
        MemoryUnit res;
        if (data->to_bool() == Logic::UNDEFINED || c.data->to_bool() == Logic::UNDEFINED) {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        } else if (data->to_bool() == Logic::FALSE || c.data->to_bool() == Logic::FALSE) {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        } else {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        }
        return res;
    }
    MemoryUnit logic_or(const MemoryUnit& c) {
        MemoryUnit res;
        if (data->to_bool() == Logic::UNDEFINED || c.data->to_bool() == Logic::UNDEFINED) {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        } else if (data->to_bool() == Logic::FALSE && c.data->to_bool() == Logic::FALSE) {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        } else {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        }
        return res;
    }
    MemoryUnit logic_nand(const MemoryUnit& c) {
        MemoryUnit res;
        if (data->to_bool() == Logic::UNDEFINED || c.data->to_bool() == Logic::UNDEFINED) {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        } else if (data->to_bool() == Logic::TRUE && c.data->to_bool() == Logic::TRUE) {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        } else {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        }
    }
    MemoryUnit logic_nor(const MemoryUnit& c) {
        MemoryUnit res;
        if (data->to_bool() == Logic::UNDEFINED || c.data->to_bool() == Logic::UNDEFINED) {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        } else if (data->to_bool() == Logic::FALSE && c.data->to_bool() == Logic::FALSE) {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        } else {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        }
    }
    MemoryUnit smaller(const MemoryUnit& c) {
        MemoryUnit res;
        if ((int)(*data) > (int)(*c.data)) {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        } else if ((int)(*data) > (int)(*c.data)) {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        } else {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        }
        return res;
    }
    MemoryUnit larger(const MemoryUnit& c) {
        MemoryUnit res;
        if ((int)(*data) > (int)(*c.data)) {
            res.data = std::make_shared<Bool>(Logic::TRUE);
        } else if ((int)(*data) < (int)(*c.data)) {
            res.data = std::make_shared<Bool>(Logic::FALSE);
        } else {
            res.data = std::make_shared<Bool>(Logic::UNDEFINED);
        }
        return res;
    }
};

class VariableUnit : public MemoryUnit {
public:
    std::string name;

    ~VariableUnit() = default;
    VariableUnit(MemoryUnit& c, std::string s) : name(std::move(s)), MemoryUnit(c) {}
    explicit VariableUnit(const std::shared_ptr<Object>& ob) : MemoryUnit(ob) {}
    [[nodiscard]] std::string get_name() const {
        return name;
    }

};

class FunctionUnit : public MemoryUnit {
public:
    std::string name;
    std::vector<std::string> params;
    std::vector<DataTypes> types;
    FuncNode* link{};

    FunctionUnit(std::string n, std::vector<std::string> names, std::vector<DataTypes> types) {
        name = std::move(n);
        for (int i = 0; i < names.size(); i++) {
            std::pair<std::string, DataTypes> p;
            p.first = names[i];
            p.second = types[i];
            params.push_back(names[i]);
            types.push_back(types[i]);
        }
    }
    ~FunctionUnit() = default;
    void set_link(FuncNode *f) {
        link = f;
    }
    FuncNode* get_link() {
        return link;
    }
    bool check_type(DataTypes t, int number) {
        for (auto i = 0; i < params.size(); i++) {
            if (types[i] == t && number == i+1) {
                return true;
            } else {
               //throw 
            }
        }
    }
};

class Memory {
private:
    std::vector<std::shared_ptr<MemoryUnit>> mem;
public:
    Memory() = default;
    ~Memory() = default;
    void add(const std::shared_ptr<MemoryUnit>& m) {
        mem.push_back(m);
    }
    void del(const std::shared_ptr<MemoryUnit>& m) {
        std::vector<MemoryUnit>::iterator it;
        int i = 0;
        for (i = 0; i < mem.size(); i++) {
            if ((*mem[i]).data == (*m).data) {
                break;
            }
        }
        if (i != -1 && i < mem.size()) {
            mem.erase(mem.begin()+i);
        }
    }
    std::shared_ptr<MemoryUnit> getVariableUnit(const std::string& s) {
        for ( auto it : mem ) {
            if (std::dynamic_pointer_cast<VariableUnit>(it)->name == s) {
                return it;
            }
        }
        return nullptr;
    }
    std::shared_ptr<MemoryUnit> getFunctionUnit(const std::string& s) {
        for ( auto it : mem ) {
            if (std::dynamic_pointer_cast<FunctionUnit>(it)->name == s) {
                return it;
            }
        }
        return nullptr;
    }
};
#endif
