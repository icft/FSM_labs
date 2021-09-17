import re
import time
import os

sample = r"^[1-9]\d*\s[a-zA-z][\w\d]*\s*=\s*(-?[1-9]\d*|[a-zA-Z][\w\d]*)(\s*[+\-\*/]\s*(-?[1-9]\d*|" \
         r"[a-zA-Z][\w\d]*))?$"

strings = ["1 b=1+",
           "1 bwdwwd=fdvmoidvnomdvslpdlwpdkwompomdpodmvvmlmvcmvlmvosdmvopdoskvddvddvdvvdvdsvvsdsdwqwdwq",
           "211212112 wopemwoemw   =    fwmdowwdwodwm        * wwinfijfwifjwfw",
           "2 fomeofmefme=1+32203898493247837483274384979372484718274812974129847812748498274724141272977428472487214",
           "67 3343oofmoemfoeclmovme                                              =      666886",
]

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
        string = str(x[i]) + ", " + str(y[i]) + "\n"
        f.write(string)


if __name__ == "__main__":
    start()
