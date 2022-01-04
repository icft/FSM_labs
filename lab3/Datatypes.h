#pragma once

#include <iostream>
#include <list>
#include <utility>
#include <vector>
#include <numeric>
#include <exception>
#include <memory>
#include "Errors.h"

enum Datatypes {INT, SHORT, BOOL, VECTOR};
enum Logic {TRUE, FALSE, UNDEFINED};

class Object {
public:
    virtual ~Object() = default;
    virtual Datatypes get_type() = 0;
    virtual bool operator==(std::shared_ptr<Object>) = 0;
    virtual void operator=(std::shared_ptr<Object>) = 0;
    virtual operator int() {
        throw CastError("Can't convert to int");
    }
    virtual operator short() {
        throw CastError("Can't convert to short");
    }
    virtual Logic to_bool() {
        throw CastError("Can't convert to bool");
    }
};

class Int : public Object {
public:
    std::int32_t value;
    Datatypes type = INT;

    Int() : value(0) {}
    Int(std::int32_t v) : value(v) {}
    Int(Int& i) : value(i.value) {}
    virtual ~Int() = default;
    virtual Datatypes get_type() {
        return type;
    }
    virtual bool operator==(std::shared_ptr<Object> obj) {
        if (value == (int)*obj) {
            return true;
        } else {
            return false;
        }
    }
    virtual void operator=(std::shared_ptr<Object> obj) {
        value = (int)*obj;
    }
    virtual operator int() {
        return value;
    }
    virtual operator short() {
        return (std::int16_t)value;
    }
    virtual Logic to_bool() {
        if (value > 0) {
            return Logic::TRUE;
        }
        else if (value < 0) {
            return Logic::FALSE;
        }
        else {
            return Logic::UNDEFINED;
        }
    }
};

class Short : public Object {
public:
    std::int16_t value;
    Datatypes type = SHORT;

    Short() : value(0) {}
    Short(std::int16_t v) : value(v) {}
    Short(Short& i) : value(i.value) {}
    virtual ~Short() = default;
    virtual Datatypes get_type() {
        return type;
    }
    virtual bool operator==(std::shared_ptr<Object> obj) {
        if (value == (int)*obj) {
            return true;
        }
        else {
            return false;
        }
    }
    virtual void operator=(std::shared_ptr<Object> obj) {
        value = (short)*obj;
    }
    virtual operator int() {
        return (std::int32_t)value;
    }
    virtual operator short() {
        return value;
    }
    virtual Logic to_bool() {
        if (value > 0) {
            return Logic::TRUE;
        }
        else if (value < 0) {
            return Logic::FALSE;
        }
        else {
            return Logic::UNDEFINED;
        }
    }
};

class Bool : public Object {
public:
    Logic value;
    Datatypes type = BOOL;

    Bool() : value(Logic::UNDEFINED) {}
    Bool(Logic v) : value(v) {}
    Bool(Bool& i) : value(i.value) {}
    virtual ~Bool() = default;
    virtual Datatypes get_type() {
        return type;
    }
    virtual bool operator==(std::shared_ptr<Object> obj) {
        if (value == obj->to_bool()) {
            return true;
        }
        else {
            return false;
        }
    }
    virtual void operator=(std::shared_ptr<Object> obj) {
        value = obj->to_bool();
    }
    virtual operator int() {
        if (value == Logic::TRUE) {
            return 1;
        }
        else if (value == Logic::FALSE) {
            return -1;
        }
        else {
            return 0;
        }
    }
    virtual operator short() {
        if (value == Logic::TRUE) {
            return 1;
        }
        else if (value == Logic::FALSE) {
            return -1;
        }
        else {
            return 0;
        }
    }
    virtual Logic to_bool() {
        return value;
    }
};


class Vector : public Object {
public:
    std::shared_ptr<Object>* vec;
    std::vector<int> dimentions;
    int count;
    Datatypes type = VECTOR;

    Vector(std::vector<int> dims) {
        auto c = 1;
        for (auto it : dims) {
            if (it < 0)
                throw IndexError("Index can't be negative");
            c *= dims[it];
        }
        //auto count = std::accumulate(std::begin(dims), std::end(dims), 1, std::multiplies<int>());
        if (c < 0) {
            throw OverflowError("Int overflow");
        }
        count = c;
        dims = dimentions;
        vec = new std::shared_ptr<Object>[count];
        /*for (int i = 0; i < count; i++) {
            vec.push_back(std::make_shared<Object>());
        }*/
    }
    Vector(std::vector<std::shared_ptr<Object>> v) {
        vec = new std::shared_ptr<Object>[v.size()];
        for (int i = 0; i < v.size(); i++)
            vec[i] = v[i];
    }
    Vector(std::vector<int> dims, std::vector<std::shared_ptr<Object>> v) {
        auto c = 1;
        for (auto it : dims) {
            if (it < 0)
                throw IndexError("Index can't be negative");
            c *= dims[it];
        }
        //auto count = std::accumulate(std::begin(dims), std::end(dims), 1, std::multiplies<int>());
        if (c < 0) {
            throw OverflowError("Int overflow");
        }
        if (c != v.size()) {
            throw SyntaxError("The dimensions must match ");
        }
        count = c;
        dimentions = dims;
        vec = new std::shared_ptr<Object>[v.size()];
        for (int i = 0; i < v.size(); i++)
            vec[i] = v[i];
    }
    Vector(Vector&& v) : vec(std::move(v.vec)), dimentions(std::move(v.dimentions)) {}
    virtual bool operator==(std::shared_ptr<Object> obj) {
        if (obj->get_type() != Datatypes::VECTOR) {
            throw TypeError("Vector cannot be compared with other types");
        }
        auto v = std::dynamic_pointer_cast<Vector>(obj);
        if (v->count != count) {
            return false;
        }
        if (v->dimentions.size() != dimentions.size()) {
            return false;
        }
        for (auto i = 0; i < dimentions.size(); i++) {
            if (v->dimentions[i] != dimentions[i]) {
                return false;
            }
        }
        for (auto i = 0; i < count; i++) {
            if (v->vec[i] != vec[i]) {
                return false;
            }
        }
        return true;
    }
    virtual void operator=(std::shared_ptr<Object> obj) {
        auto v = std::dynamic_pointer_cast<Vector>(obj);
        vec = v->vec;
        dimentions = v->dimentions;
    }
    virtual Datatypes get_type() {
        return type;
    }
    std::shared_ptr<Object> operator[](std::vector<int> dims) {
        if (dimentions.size() != dims.size()) {
            throw OverflowError("Out of range");
        }
        std::vector<int> s;
        for (auto i = 0; i < dimentions.size(); i++) {
            s.push_back(1);
            for (auto j = i + 1; j < dimentions.size(); j++) {
                s[i] *= dimentions[j];
            }
        }
        int index = 0;
        for (auto i = 0; i < s.size(); i++) {
            if (dims[i] < 0)
                throw IndexError("Index can't be negative");
            if (dims[i]+1 > dimentions[i]) {
                throw OverflowError("Out of range");
            }
            index += s[i] * dims[i];
        }
        return vec[index];
    }
    virtual ~Vector() = default;
};
