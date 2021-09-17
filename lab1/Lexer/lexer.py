import re
import lex_yacc.lex as lex

d = {}

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

        # try:
        #     print(lexer.lexdata)
        #     tok = lexer.token()
        #     num = int(re.findall(r'\w+', string)[0])
        #     try:
        #         d[num] += 1
        #     except KeyError:
        #         d[num] = 1
        # except lex.LexError:
        #     continue


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