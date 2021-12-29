#pragma once
#include <iostream>
#include <exception>
#include <string>

class SyntaxError : public std::exception {
public:
	SyntaxError(std::string s) : exception(s.c_str()) {}
};

class CastError : public std::exception {
public:
	CastError(std::string s) : exception(s.c_str()) {}
};

class IndexError : public std::exception {
public:
	IndexError(std::string s) : exception(s.c_str()) {}
};

class NameError : public std::exception {
public:
	NameError(std::string s) : exception(s.c_str()) {}
};

class OverflowError : public std::exception {
public:
	OverflowError(std::string s) : exception(s.c_str()) {}
};

class TypeError : public std::exception {
public:
	TypeError(std::string s) : exception(s.c_str()) {}
};

class MemoryError : public std::exception {
public:
	MemoryError(std::string s) : exception(s.c_str()) {}
};
