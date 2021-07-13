resp = ""
homem = 0
maiores18 = 0
fem20 = 0

print("Analise de Grupo")
print("=+="*6)
while resp != "N":
    idade = int(input("Digite a idade:"))
    sexo = ""
    while sexo != "M" and sexo != "F":
        sexo = str(input("Digite o sexo[M]/[F]:")).upper() 
    print("=+="*6)
    if sexo == "M":
        homem += 1
    if idade >= 18:
        maiores18 += 1
    if sexo == "F" and idade < 20:
        fem20 += 1
    resp = ""
    while resp != "S"and resp != "N":
        resp = str(input("Quer registrar mais um[S]/[N]:")).upper()
    print("=+="*6)

print("FIM DO REGISTRO")
print("=+="*6)
print(" "*3,"RESULTADOS")
print(f"Maiores de 18 = {maiores18}")
print(f"Homens cadastrado = {homem}")
print(f"Mulheres com menos de 20 anos = {fem20}")
