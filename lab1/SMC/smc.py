import fsm_sm
import fsm

d = {}


def check(a):
    for j in range(a):
        string = input()
        machine = fsm.Fsm()
        flag, num = machine.parse(string)
        if flag:
            try:
                d[num] += 1
            except KeyError:
                d[num] = 1


if __name__ == "__main__":
    count = int(input("Enter the number of strings: "))
    print("Enter strings:")
    check(count)
    print("Correct string usage statistics:")
    if len(d) == 0:
        print("Empty")
    else:
        for key in sorted(d.keys()):
            print("{0}: {1}".format(key, d[key]))