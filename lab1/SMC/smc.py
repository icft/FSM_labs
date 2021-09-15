import fsm_sm
import fsm

d = {}
true = set()
false = set()


def check(a):
    for j in range(a):
        string = input()
        machine = fsm.Fsm()
        flag, num = machine.parse(string)
        if flag:
            true.add(num)
            try:
                d[num] += 1
            except KeyError:
                d[num] = 1
        else:
            false.add(num)


if __name__ == "__main__":
    count = int(input("Enter the number of strings: "))
    print("Enter strings:")
    check(count)
    print("Correct strings:")
    if len(true) == 0:
        print("Empty")
    else:
        print(*sorted(list(true)))
    print("Incorrect strings:")
    if len(false) == 0:
        print("Empty")
    else:
        print(*sorted(list(false)))
    print("Correct string usage statistics:")
    for key in sorted(d.keys()):
        print("{0}: {1}".format(key, d[key]))