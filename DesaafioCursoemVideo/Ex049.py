n = int(input("Qual tabuada voce quer:"))
c = 0

for c in range(0,11):
    print("{} x {} = {}".format(n, c, n*c))
    c += 1
