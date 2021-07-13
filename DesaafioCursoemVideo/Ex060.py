number =  int(input("Digite um valor:"))
c = 0
cache = 1
print(f"Calculando... {number}! =", end="")
for num in range(number, c, -1):
        print(f" {num}",end="")
        print(f" x " if num > 1 else " =",end=" " )
        cache *= num
print(cache)


