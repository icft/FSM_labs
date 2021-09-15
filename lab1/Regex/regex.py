import re

sample = r"[1-9]\d*\s[a-zA-z]+\s*=\s*" \
         r"((((-[1-9]|[1-9])\d*)|[a-zA-Z]+)\s*[+\-/\*]\s*(((-[1-9]|[1-9])\d*)|[a-zA-Z]+)" \
         r"|(((-[1-9]|[1-9])\d*)|[a-zA-Z]+))"

d = {}
true = set()
false = set()


#if re.compile(sample).match(str):


def check(a):
    for j in range(a):
        string = input()
        num = int(re.findall(r'\w+', string)[0])
        pattern = re.compile(sample)
        if pattern.match(string):
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