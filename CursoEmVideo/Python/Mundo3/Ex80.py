# Exercício Python 080: 
# Crie um programa onde o usuário possa digitar cinco valores numéricos e cadastre-os em uma lista, já na posição correta de inserção (sem usar o sort()).
#  No final, mostre a lista ordenada na tela.

lista: list = []

for i in range(1,6):
    numInput: int = int(input("Digite um valor:"))
    if len(lista) > 0:
        for num in lista:
            if numInput < num: # se numInput for menor que o primeiro elemento da lista
                lista.insert(lista.index(num),numInput)
                break
            elif numInput > num: # se numInput for maior que o último elemento da lista
                lista.append(numInput)
                break
    else:
        lista.append(numInput)

print(lista)