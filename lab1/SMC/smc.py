import fsm


def check(a, machine):
    for j in range(a):
        string = input()
        machine.parse(string)


def start():
    machine = fsm.Fsm()
    count = int(input("Enter the number of strings: "))
    print("Enter strings:")
    check(count, machine)
    print("Correct string usage statistics:")
    machine.displayStatistics()


if __name__ == "__main__":
    start()