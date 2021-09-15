import re

sample = r"^[1-9]\d*\s[a-zA-z]+\s*=\s*(-?\d+|\w[\w\d]*)(\s*[+\-\*/]\s*(-?\d+|\w[\w\d]*))?$"
d = {}


def check(a):
    for j in range(a):
        string = input()
        num = None
        try:
            num = int(re.findall(r'\w+', string)[0])
        except ValueError:
            pass
        pattern = re.compile(sample)
        if pattern.match(string):
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