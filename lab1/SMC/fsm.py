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
        self.fsm.enterStartState()

    def parse(self, string):
        s = ""
        num = None
        for c in string:
            b = self.fsm.getState()
            if b == fsm_sm.FSM.q0 or b == fsm_sm.FSM.q1:
                s += c
            if b == fsm_sm.FSM.error:
                break
            if c in digit:
                if b == fsm_sm.FSM.q0 or b == fsm_sm.FSM.q5 or \
                        b == fsm_sm.FSM.q6 or b == fsm_sm.FSM.q7:
                    try:
                        self.fsm.natural()
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
                else:
                    try:
                        self.fsm.digit()
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
            elif 97 <= ord(c) <= 122 or 65 <= ord(c) <= 90:
                try:
                    self.fsm.alpha()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            elif c == separator:
                try:
                    self.fsm.separator()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            elif c == equal:
                try:
                    self.fsm.equal()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            elif c in operations:
                try:
                    self.fsm.operations()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            else:
                self.fsm.err()
        if self.fsm.getState() == fsm_sm.FSM.error:
            self.flag = False
        else:
            try:
                self.fsm.EOS()
                self.flag = True
                num = int(s)
            except statemap.TransitionUndefinedException:
                self.fsm.err()
        self.fsm.setState(fsm_sm.FSM.q0)
        return self.flag, num
