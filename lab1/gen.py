import random
import exrex

sample = r"^[1-9]\d* [a-zA-z][a-zA-Z\d]{,15} = (-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})( [+\-\*/] (-?[1-9]\d*|" \
         r"[a-zA-Z][a-zA-Z\d]{,15}))?$"

errSample1 = r"^[1-9]\d* [a-zA-z][a-zA-Z\d]{,15} = (-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{,15})[+\-\*/]$"
errSample2 = r"^[1-9]\d* [a-zA-z][a-zA-Z\d]{,15} = $"
errSample3 = r"^[1-9]\d*$"
errSample4 = r"^[1-9]\d* [a-zA-z][a-zA-Z\d]{,15} = (-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{16,30})"
errSample5 = r"^[1-9]\d* [a-zA-z][a-zA-Z\d]{16,30} = (-?[1-9]\d*|[a-zA-Z][a-zA-Z\d]{16,30})"


samples = [sample, errSample1, errSample2, errSample3, errSample4]
f = open("strings.txt", "w")

for i in range(100000):
    a = random.randint(0, 4)
    s = exrex.getone(samples[a])
    f.write(s+'\n')
