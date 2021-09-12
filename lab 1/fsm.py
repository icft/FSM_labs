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

    def __int__(self):
        self.fsm = fsm_sm.Fsm_sm(self)
        self.flag = False

    def Parse(self, string):
        pass