import fsm_sm

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
        for c in string:
            if c in natural:
                self.fsm.natural()
            if c in digit:
                self.fsm.digit()
            if 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90:
                self.fsm.alpha()
            if 97 <= ord(c) <= 122 and c in digit or 65 <= ord(c) <= 90 and c in digit:
                self.fsm.alnum()
            if c == separator:
                self.fsm.separator()
            if c == equal:
                self.fsm.equal()
            if c in operations:
                self.fsm.operations()
            else:
                break
        return self.flag
