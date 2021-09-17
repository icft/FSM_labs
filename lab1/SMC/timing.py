import fsm
import time


strings = ["1 b=1+",
           "1 bwdwwd=fdvmoidvnomdvslpdlwpdkwompomdpodmvvmlmvcmvlmvosdmvopdoskvddvddvdvvdvdsvvsdsdwqwdwq",
           "211212112 wopemwoemw   =    fwmdowwdwodwm        * wwinfijfwifjwfw",
           "2 fomeofmefme=1+32203898493247837483274384979372484718274812974129847812748498274724141272977428472487214",
           "67 3343oofmoemfoeclmovme                                              =      666886",
]

def timing(x, y):
    machine = fsm.Fsm()
    for i in strings:
        machine = fsm.Fsm()
        start_time = time.perf_counter()
        flag, num = machine.parse(i)
        y.append(time.perf_counter() - start_time)
        x.append(len(i))


def start():
    x, y = [], []
    timing(x, y)
    f = open('smc.txt', 'w+')
    for i in range(len(x)):
        string = str(x[i]) + ", " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()
