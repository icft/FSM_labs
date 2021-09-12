import fsm_sm
import fsm

d = {}


def check():
    while True:
        string = input()
        machine = fsm.Fsm()
        flag, num = machine.parse(string)
        try:
            d[num] += 1
        except KeyError:
            d[num] = 1


#if __name__ == "__main__":
#    check()
#    for key, values in d.items():
#        print("{0}: {1}".format(key, values))

for property, value in vars(fsm_sm.FSM_q0).iteritems():
    print(property, ":", value)