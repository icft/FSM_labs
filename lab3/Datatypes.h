#pragma once
#ifndef INTERPRETER_DATATYPES_H
#define INTERPRETER_DATATYPES_H

#include <iostream>
#include <list>
#include <utility>
#include <vector>
#include <numeric>

enum class DataTypes { INT, SHORT, BOOL, VECTOR };
enum class Logic { TRUE, FALSE, UNDEFINED };
class Object {
public:
    virtual ~Object() = 0;
    [[nodiscard]] virtual DataTypes get_type() const noexcept = 0;
    virtual std::shared_ptr<Object> copy() = 0;
    virtual bool operator==(std::shared_ptr<Object>) = 0;
    virtual explicit operator int() = 0; //написать ошибку
    virtual Logic to_bool() = 0; //написать ошибку
};

class Int : public Object {
private:
    int value;
    DataTypes type = DataTypes::INT;
public:
    Int() : value(0) {}
    explicit Int(int val) : value(val) {}
    Int(Int& i) : value(i.value) {}
    std::shared_ptr<Object> copy() override {
        return std::make_shared<Int>(value);
    }
    [[nodiscard]] DataTypes get_type() {
        return type;
    }
    bool operator==(Int& i) const {
        if (value == i.value) {
            return true;
        }
        return false;
    }
    explicit operator int() override {
        return value;
    }
    Logic to_bool() override {
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
    ~Int() override = default;
};

class Short : public Object {
private:
    short value;
    DataTypes type = DataTypes::SHORT;
public:
    Short() : value(0) {}
    explicit Short(short val) : value(val) {}
    Short(Short& s) : value(s.value) {}
    Short(Short&& s) noexcept : value(s.value) {}
    std::shared_ptr<Object> copy() override {
        return std::make_shared<Short>(value);
    }
    [[nodiscard]] DataTypes get_type() {
        return type;
    }
    bool operator==(Short& s) const {
        if (value == s.value) {
            return true;
        }
        return false;
    }
    explicit operator int() override {
        return value;
    }
    Logic to_bool() override {
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
    ~Short() override = default;
};

class Bool : public Object {
private:
    Logic value;
    DataTypes type = DataTypes::BOOL;
public:
    Bool() : value(Logic::UNDEFINED) {}
    explicit Bool(Logic val) : value(val) {}
    Bool(Bool& b) : value(b.value) {}
    Bool(Bool&& b) noexcept : value(b.value) {}
    std::shared_ptr<Object> copy() override {
        return std::make_shared<Bool>(value);
    }
    [[nodiscard]] DataTypes get_type() {
        return type;
    }
    bool operator==(Bool& b) const {
        if (value == b.value) {
            return true;
        }
        return false;
    }
    explicit operator int() override {
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
    Logic to_bool() override {
        return value;
    }
    ~Bool() override = default;
};

class Vector : public Object {
private:
    std::vector<std::shared_ptr<Object>> vec;
    std::vector<int> dim;
    DataTypes type = DataTypes::VECTOR;
public:
    Vector() = default;
    Vector(std::initializer_list<int> dims) {
        auto count = std::accumulate(std::begin(dims), std::end(dims), 1, std::multiplies());
        if (count < 0) {
            throw std::overflow_error("Overflow");
        }
        for (int i = 0; i < count; i++) {
            vec.push_back(std::make_shared<Object>());
        }
        dim = dims;
    }
    explicit Vector(const std::vector<std::shared_ptr<Object>>& v) {
        for (auto i : v) {
            vec.push_back(i);
        }
    }
    explicit Vector(const std::vector<std::shared_ptr<Object>>& v, const std::vector<int>& dims) {
        dim = dims;
        for (auto i : v) {
            vec.push_back(i);
        }

    }
    Vector(Vector& v) {
        for (auto i : v.vec) {
            vec.push_back(i);
        }
        dim = v.dim;
    }
    Vector(Vector&& v)  noexcept : vec(std::move(v.vec)), dim(std::move(v.dim)) {}
    std::shared_ptr<Object> copy() override {
        return std::make_shared<Vector>(vec);
    }
    bool operator==(Vector& v) const {
        if (v.vec.size() != vec.size()) {
            return false;
        }
        if (v.dim.size() != dim.size()) {
            return false;
        }
        for (auto i = 0; i < dim.size(); i++) {
            if (!(v.dim[i] == v.dim[i])) {
                return false;
            }
        }
        for (auto i = 0; i < vec.size(); i++) {
            if (!(v.vec[i] == v.vec[i])) {
                return false;
            }
        }
        return true;
    }
    [[nodiscard]] DataTypes get_type() {
        return type;
    }
    std::shared_ptr<Object> operator[](std::vector<int> dims) {
        std::vector<int> s;
        for (auto i = 0; i < dim.size(); i++) {
            s.push_back(1);
            for (auto j = i + 1; j < dim.size(); j++) {
                s[i] *= dim[j];
            }
        }
        int index = 0;
        for (auto i = 0; i < s.size(); i++) {
            index += s[i] * dims[i];
        }
        return vec[index];
    }
    ~Vector() override = default;
};
#endif
