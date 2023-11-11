# Exercício Python 084: Faça um programa que leia nome e peso de várias pessoas,
# guardando tudo em uma lista. No final mostre:
# A) Quantas pessoas foram cadastradas.              
# B) Uma listagem com as pessoas mais pesadas.
# C) Uma listagem com as pessoas mais leves.

listaPeso: list = []
listaNome: list = []
lista = []

quantidadePessoas: int = int(input('Quantas pessoas você quer cadastrar? '))

for i in range(0, quantidadePessoas):
    nome: str = str(input('Nome: '))
    peso: float = float(input('Peso: '))
    listaNome.append(nome)
    listaPeso.append(peso)

print(f'Foram cadastradas {quantidadePessoas} pessoas.')
print(f'O maior peso foi {max(listaPeso)}Kg. Peso de ', end='')
for i in listaPeso:
    if listaPeso[i] == max(listaPeso):
        print(f'{listaNome[i]} ', end='')
print(f'\nO menor peso foi {min(listaPeso)}Kg. Peso de ', end='')
for i in listaPeso:
    if listaPeso[i] == min(listaPeso):
        print(f'{listaNome[i]} ', end='')