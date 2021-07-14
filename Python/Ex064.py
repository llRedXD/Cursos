number = 0
soma = 0
while number != 999:
    number = int(input("Digite um numero:"))
    if number != 999:
        soma += number
print(f"A soma dos valores digitados é de {soma}")
