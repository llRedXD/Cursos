number = int(input("Digite um valor:"))
c = 0
cache = 1
print(f"Calculando... {number}! =", end="")
while c < number:
    print(f" {number}", end="")
    print(f" x " if number > 1 else " =", end=" ")
    cache *= number
    number -= 1
print(cache)
