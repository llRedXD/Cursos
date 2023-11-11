# Exercício Python 082: Crie um programa que vai ler vários números e colocar em uma lista.
# Depois disso, crie duas listas extras que vão conter apenas os valores pares e os valores ímpares digitados, respectivamente. Ao final, mostre o conteúdo das três listas geradas.

listaPar = []
listaImpar = []
lista = []

quantDesejada = int(input('Quantos números deseja digitar? '))

for i in range(0, quantDesejada):
    lista.append(int(input(f'Digite o {i+1}º número: ')))

for i in lista:
    if i % 2 == 0:
        listaPar.append(i)
    if i % 2 == 1:
        listaImpar.append(i)

print(f'Lista completa: {lista}')
print(f'Lista par: {listaPar}')
print(f'Lista ímpar: {listaImpar}')