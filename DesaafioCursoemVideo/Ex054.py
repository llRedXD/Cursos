c = 0 
maior = 0
menor = 0
for c in range(0,7):
    nascimento = int(input("Qual o ano de nascimento:"))
    idade = 2021 - nascimento
    if idade <= 18:
        maior += 1
    else:
        menor += 1

print(f"O numero de maiores de idade é {maior} e de menores {menor}")