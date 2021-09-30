import fsm
import time


name = "./strings.txt"
file = open(name, "r")
strings = file.readlines()


def timing(x, y):
    machine = fsm.Fsm()
    for s in strings:
        start_time = time.perf_counter()
        machine.parse(s.strip())
        y.append(time.perf_counter() - start_time)
        x.append(len(s))


def start():
    x, y = [], []
    timing(x, y)
    f = open('smc.txt', 'w+')
    for i in range(len(x)):
        string = str(x[i]) + " " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()
