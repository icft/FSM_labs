import fsm_sm
import statemap

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
        self.flag = None

    def parse(self, string):
        self.fsm.enterStartState()
        s = ""
        i = 0
        for c in string:
            if c != " " and i == 0:
                s += c
            if c == " ":
                i += 1
            b = self.fsm.getState()
            if c in digit:
                if b.getId() == 0 or b.getId() == 5 or b.getId() == 6 or b.getId() == 7:
                    self.fsm.natural()
                else:
                    self.fsm.digit()
            elif 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90:
                self.fsm.alpha()
            elif c == separator:
                self.fsm.separator()
            elif c == equal:
                self.fsm.equal()
            elif c in operations:
                self.fsm.operations()
            else:
                return False, int(s)
        try:
            self.fsm.EOS()
            self.flag = True
        except statemap.TransitionUndefinedException:
            self.flag = False
        return self.flag, int(s)