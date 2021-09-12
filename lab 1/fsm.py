import fsm_sm
import statemap

natural = "123456789"
digit = "0123456789"
separator = " "
ascii_lower_start = 97
ascii_lower_end = 122
ascii_upper_start = 65
ascii_upper_end = 90
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
            if c in natural:
                self.fsm.natural()
            if c in digit:
                self.fsm.digit()
            if 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90:
                self.fsm.alpha()
            if 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90 or c in digit:
                self.fsm.alnum()
            if c == separator:
                self.fsm.separator()
            if c == equal:
                self.fsm.equal()
            if c in operations:
                self.fsm.operations()
            else:
                break

        try:
            self.fsm.EOS()
        except statemap.TransitionUndefinedException:
            self.flag = False

        return self.flag, i
