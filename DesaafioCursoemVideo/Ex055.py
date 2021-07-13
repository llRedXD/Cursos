c = 0
peso = 1 
maior_peso = 0
menor_peso = 0
for c in range(0,5):
    peso = float(input("Digite o peso:"))
    if c == 1:
        maior_peso = peso
        menor_peso = peso
    else:
        if peso > maior_peso:
            maior_peso = peso
        elif menor_peso >= peso:
            menor_peso = peso
print(f"O maior peso foi de {maior_peso}Kg e o menor de {menor_peso}Kg") 
