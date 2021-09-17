import re
import time

import lex_yacc.lex as lex

d = {}
strings = ["1 b=1+",
           "1 bwdwwd=fdvmoidvnomdvslpdlwpdkwompomdpodmvvmlmvcmvlmvosdmvopdoskvddvddvdvvdvdsvvsdsdwqwdwq",
           "211212112 wopemwoemw   =    fwmdowwdwodwm        * wwinfijfwifjwfw",
           "2 fomeofmefme=1+32203898493247837483274384979372484718274812974129847812748498274724141272977428472487214",
           "67 3343oofmoemfoeclmovme                                              =      666886",
]


tokens = ("NUM", "STR")
t_NUM = r"[1-9]\d*"
t_STR = r"[a-zA-z][a-zA-Z\d]*\s*=\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]*)(\s*[+\-\*/]\s*(-?[1-9]\d*|" \
        r"[a-zA-Z][a-zA-Z\d]*))?"
t_ignore = " \r\n\t\f"


def t_error(t):
    # print("Illegal character %s" % t.value[0])
    # t.lexer.skip(1)
    pass


lexer = lex.lex()

def timing(x, y):
    for s in strings:
        lexer.input(s)
        start_time = time.perf_counter()
        try:
            tok1 = lexer.token()
            tok2 = lexer.token()
            end_time = time.perf_counter()
            if tok1 and tok2:
                y.append(end_time-start_time)
                x.append(len(s))
        except lex.LexError:
            end_time = time.perf_counter()
            y.append(end_time-start_time)
            x.append(len(s))



def check(a):
    for j in range(a):
        string = input()
        lexer.input(string)
        try:
            tok = lexer.token()
            if tok:
                num = tok.value
                try:
                    tok = lexer.token()
                    try:
                        d[num] += 1
                    except KeyError:
                        d[num] = 1
                except lex.LexError:
                    continue
        except lex.LexError:
            continue


def start():
    x, y = [], []
    timing(x, y)
    f = open('lexer.txt', 'w+')
    for i in range(len(x)):
        string = str(x[i]) + ", " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()