add = 0
c = 0
for c in range(0,501):
    if c%2 == 1:
        if c%3 == 0:
            add = add + c
            print(f"\033[31;40m{c}|\033[m", end="")
print(f"\nA SOMA DOS NUMEROS: \033[31;40m{add}\033[m")