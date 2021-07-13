nome_maior = ""
mais_velho = 0
idade = 0
mulher = 0
soma = 0
c = 0
for c in range(0,4):
    nome = str(input("Digite o nome:"))
    idade = int(input("Digite a idade:"))
    sexo = int(input("[1] Masculino\n[2] Feminino\n"))
    if sexo == 1:
        sexo = "Masculino"
    else:
        sexo = "Feminino"
    if idade > mais_velho and sexo == "Masculino":
        mais_velho = idade
        nome_maior = nome
    if sexo == "Feminino":
        mulher += 1
    soma = soma + idade
media = soma / 4
print(f"A media de idade do grupo é de {media:.0f}")
print(f"O homem mais velho é {nome_maior} tendo {mais_velho} anos ")
print(f"O grupo tem {mulher} Mulheres")
