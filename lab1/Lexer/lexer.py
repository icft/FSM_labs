import re
import lex_yacc.lex as lex

d = {}

# tokens = ("NUM", "STR")
tokens = ("NUM", "LPART", "RPART")
t_NUM = r"[1-9]\d*"
# t_STR = r"[a-zA-z][a-zA-Z\d]{,15}\s*=\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})(\s*[+\-\*/]\s*(-?[1-9]\d*|" \
#         r"[a-zA-Z][a-zA-Z\d]{,15}))?"
t_LPART = r"[a-zA-z][a-zA-Z\d]{,15}\s*=\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})"
t_RPART = r"\s*[+\-\*/]\s*(-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})"
t_ignore = " \r\n\t\f"


def t_error(t):
    # print("Illegal character %s" % t.value[0])
    # t.lexer.skip(1)
    pass


lexer = lex.lex()


# def check(a):
#     for j in range(a):
#         string = input()
#         lexer.input(string)
#         try:
#             tok1 = lexer.token()
#             tok2 = lexer.token()
#             if tok1 and tok2:
#                 num = tok1.value
#                 try:
#                     d[num] += 1
#                 except KeyError:
#                     d[num] = 1
#         except lex.LexError:
#             continue


def check(a):
    for j in range(a):
        b = []
        string = input()
        lexer.input(string)
        while True:
            try:
                tok = lexer.token()
                if not tok:
                    break
                b.append(tok)
            except lex.LexError:
                break
#        print(b)
        f1 = len(b) == 3 and b[0].type == "NUM" and\
             b[1].type == "LPART" and b[2].type == "RPART"
        f2 = len(b) == 2 and b[0].type == "NUM" and b[1].type == "LPART"
        if f1 or f2:
            try:
                d[int(a[0].value)] += 1
            except KeyError:
                d[int(a[0].value)] = 1


def start():
    count = int(input("Enter the number of strings: "))
    print("Enter strings:")
    check(count)
    print("Correct string usage statistics:")
    if len(d) == 0:
        print("Empty")
    else:
        for key in sorted(d.keys()):
            print("{0}: {1}".format(key, d[key]))


if __name__ == "__main__":
    start()