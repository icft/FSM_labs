import re
import time

import lex_yacc.lex as lex

d = {}
name = "./strings.txt"
file = open(name, "r")
strings = file.readlines()


tokens = ("NUM", "LPART", "RPART")
t_NUM = r"[1-9]\d*"
#t_STR = r"[a-zA-z][a-zA-Z\d]{,15}\s*=\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})(\s*[+\-\*/]\s*(-?[1-9]\d*|" \
#        r"[a-zA-Z][a-zA-Z\d]{,15}))?"
t_LPART = r"[a-zA-z][a-zA-Z\d]{,15}\s*=\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})"
t_RPART = r"\s*[+\-\*/]\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})"
t_ignore = " \r\n\t\f"


def t_error(t):
    # print("Illegal character %s" % t.value[0])
    # t.lexer.skip(1)
    pass


lexer = lex.lex()


def timing(x, y):
    for s in strings:
        a = []
        lexer.input(s)
        start_time = time.perf_counter()
        try:
            while True:
                try:
                    tok = lexer.token()
                    if not tok:
                        break
                    a.append(tok)
                except lex.LexError:
                    break
            if len(a) == 3 and a[0].type == "NUM" and a[1].type == "LPART" and a[2].type == "RPART":
                end_time = time.perf_counter()
                y.append(end_time - start_time)
                x.append(len(s))
            elif len(a) == 2 and a[0].type == "NUM" and a[1].type == "LPART":
                end_time = time.perf_counter()
                y.append(end_time - start_time)
                x.append(len(s))
        except lex.LexError:
            end_time = time.perf_counter()
            y.append(end_time-start_time)
            x.append(len(s))



def start():
    x, y = [], []
    timing(x, y)
    f = open('lexer.txt', 'w+')
    for i in range(len(x)):
        string = str(x[i]) + " " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()