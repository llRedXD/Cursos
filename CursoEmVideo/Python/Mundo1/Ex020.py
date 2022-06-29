import random
nome = []
print("Ordem de apresentação")
print("="*27)
c = 0
while c < 4:
    nome.append(input("Nome:"))
    c += 1
print("O escolhido é {}\n".format(random.sample(nome,k=4)))
