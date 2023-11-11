# Exercício Python 081: Crie um programa que vai ler vários números e colocar em uma lista.
# Depois disso mostre:              
#   A) Quantos números foram.
#   B) A lista de valores, ordenada de forma decrescente. 
#   C) Se o valor 5 foi digitado e está ou não na lista.


tamanhoLista: int = int(input("Digite o tamanho da lista: "))
lista: list = []

for i in range(0, tamanhoLista):
    lista.append(int(input("Digite um valor: ")))

print(f"Quantidade de números digitados: {len(lista)}")
print(f"Lista ordenada de forma decrescente: {sorted(lista, reverse=True)}")
if 5 in lista:
    print("O valor 5 está na lista")