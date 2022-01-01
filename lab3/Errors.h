#pragma once
#include <iostream>
#include <exception>
#include <string>
#include <utility>


class Exception : public std::exception {
public:
    std::string msg;

    Exception(std::string s) : msg(std::move(s)) {}
    virtual void what() {std::cout << msg << std::endl;}
    virtual ~Exception() = default;
};

class SyntaxError : public Exception {
public:
    SyntaxError(std::string s) : Exception(std::move(s)) {}
    virtual ~SyntaxError() = default;
};

class CastError : public Exception {
public:
    CastError(std::string s) : Exception(std::move(s)) {}
    virtual ~CastError() = default;
};

class IndexError : public Exception {
public:
    IndexError(std::string s) : Exception(std::move(s)) {}
    virtual ~IndexError() = default;
};

class NameError : public Exception {
public:
    NameError(std::string s) : Exception(std::move(s)) {}
    virtual ~NameError() = default;
};

class OverflowError : public Exception {
public:
    OverflowError(std::string s) : Exception(std::move(s)) {}
    virtual ~OverflowError() = default;
};

class TypeError : public Exception {
public:
    TypeError(std::string s) : Exception(std::move(s)) {}
    virtual ~TypeError() = default;
};

class MemoryError : public Exception {
public:
    MemoryError(std::string s) : Exception(std::move(s)) {}
    virtual ~MemoryError() = default;
};
