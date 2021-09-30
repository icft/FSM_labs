import re
import time
import os

sample = r"^[1-9]\d*\s[a-zA-z][\w\d]*\s*=\s*(-?[1-9]\d*|[a-zA-Z][\w\d]*)(\s*[+\-\*/]\s*(-?[1-9]\d*|" \
         r"[a-zA-Z][\w\d]*))?$"

name = "./strings.txt"
file = open(name, "r")
strings = file.readlines()


def timing(x, y):
    pattern = re.compile(sample)
    for s in strings:
        start_time = time.perf_counter()
        pattern.match(s)
        y.append(time.perf_counter() - start_time)
        x.append(len(s))


def start():
    x, y = [], []
    timing(x, y)
    f = open('regex.txt', 'w+')
    for i in range(len(x)):
        string = str(x[i]) + " " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()
