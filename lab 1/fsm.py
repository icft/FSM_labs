import fsm_sm
import statemap

natural = "123456789"
digit = "0123456789"
separator = " "
ascii_lower_start = 97
ascii_lower_end = 122
ascii_upper_start = 65
ascii_upper_end = 90
minus = "-"
equal = "="
operations = "+-*/"


class Fsm:

    def __init__(self):
        self.fsm = fsm_sm.Fsm_sm(self)
        self.flag = False

    def parse(self, string):
        self.fsm.enterStartState()
        s = ""
        i = 0
        for c in string:
            if c != " " and i == 0:
                s += c
            if c == " ":
                i += 1
            print(c)
            print(self.fsm.getState())
            if c in natural:
                self.fsm.natural()
            elif c in digit:
                self.fsm.digit()
            elif 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90:
                self.fsm.alpha()
            elif 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90 or c in digit:
                self.fsm.alnum()
            elif c == minus:
                self.fsm.minus()
            elif c == separator:
                self.fsm.separator()
            elif c == equal:
                self.fsm.equal()
            elif c in operations:
                self.fsm.operations()
            else:
                break
        try:
            self.fsm.EOS()
            self.flag = True
        except statemap.TransitionUndefinedException:
            self.flag = False
        print(self.flag)
        return self.flag, int(s)