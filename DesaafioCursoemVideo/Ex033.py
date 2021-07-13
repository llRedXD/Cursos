cont = 0
Mnun = 0
while cont <= 2:
    num = int(input("Digite um numero:"))
    if num > Mnun:
        Mnun = num
    cont += 1
print("O maior numero é {}".format(Mnun))