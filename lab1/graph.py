import matplotlib.pyplot as plt

names = ["Lexer/lexer.txt", "Regex/regex.txt", "SMC/smc.txt"]
colors = ['r', 'g', 'b']
legend = ["ply", "regex", "smc"]


def read(d, name):
    with open(name, 'r') as f:
        for line in f:
            res = line.split()
            d[int(res[0])] = float(res[1])


def start():
    plt.title("Timing")
    plt.xlabel("String length")
    plt.ylabel("Time")
    d = {}
    for i in range(3):
        read(d, names[i])
        d = dict(sorted(d.items(), key=lambda x: x[0]))
        plt.plot(d.keys(), d.values(), label=legend[i])
        d.clear()
    plt.legend()
    plt.savefig("comparison")


if __name__ == "__main__":
    start()