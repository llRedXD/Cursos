escolha = ""
soma = 0
maiornum = 0
menornum = 1
vezes = 0
while escolha != "N":
    number = int(input("Digite um valor:"))
    if escolha == "":
        maiornum = number
        menornum = number
    else:
        if number > maiornum:
            maiornum = number
        elif number < menornum:
            menornum = number
    escolha = str(input("Quer digitar mais valores? [S]/[N]\n :")).upper()
    soma = soma + number
    if "S" and "N":
        vezes += 1

media = soma / vezes
print(f"O maior numero digitado foi {maiornum} e o menor {menornum}.\n A media dos numero foi :{media}")
