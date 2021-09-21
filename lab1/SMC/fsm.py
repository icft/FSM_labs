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
        self.d = dict()
        self.s = ""
        self.count = 0
        self.fsm.enterStartState()

    def createNumber(self, c):
        self.s += c

    def isValid(self):
        if self.count <= 16:
            return True
        else:
            return False

    def resetStr(self):
        self.s = ""

    def resetCounter(self):
        self.count = 0

    def increase(self):
        self.count += 1

    def parse(self, string):
        for c in string:
            b = self.fsm.getState()
            if b == fsm_sm.FSM.error:
                break
            if c in digit:
                if b == fsm_sm.FSM.q0:
                    try:
                        self.fsm.natural_create(c)
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
                elif b == fsm_sm.FSM.q1:
                    try:
                        self.fsm.digit_create(c)
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
                elif b == fsm_sm.FSM.q3 or b == fsm_sm.FSM.q8:
                    try:
                        self.fsm.digit_with_check()
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
                elif b == fsm_sm.FSM.q5 or b == fsm_sm.FSM.q6 or b == fsm_sm.FSM.q7:
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
                if b == fsm_sm.FSM.q3 or b == fsm_sm.FSM.q8:
                    try:
                        self.fsm.alpha_with_check()
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
                else:
                    try:
                        self.fsm.alpha()
                    except statemap.TransitionUndefinedException:
                        self.fsm.err()
            elif c == separator:
                self.resetCounter()
                try:
                    self.fsm.separator()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            elif c == equal:
                self.resetCounter()
                try:
                    self.fsm.equal()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            elif c in operations:
                self.resetCounter()
                try:
                    self.fsm.operations()
                except statemap.TransitionUndefinedException:
                    self.fsm.err()
            else:
                self.fsm.err()
        if self.fsm.getState() != fsm_sm.FSM.error:
            try:
                self.fsm.EOS()
                try:
                    self.d[int(self.s)] += 1
                except KeyError:
                    self.d[int(self.s)] = 1
            except statemap.TransitionUndefinedException:
                self.fsm.err()
        self.fsm.setState(fsm_sm.FSM.q0)
        self.resetCounter()
        self.resetStr()

    def displayStatistics(self):
        if len(self.d) == 0:
            print("Empty")
        else:
            for key in sorted(self.d.keys()):
                print("{0}: {1}".format(key, self.d[key]))