n = int(input("Digite um valor:"))
mult = 0
for cont in range(2,n):
    if n%cont == 0:
        print(f"Multiplo de {cont}")
        mult += 1
if mult == 0:
    print(f"{n} é primo")
else:
    print(f"{n} nao é primo")