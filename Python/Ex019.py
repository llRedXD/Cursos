import random
nome = []
print("Sorteio para apagar a lousa")
print("="*27)
c = 0
while c < 4:
    nome.append(input("Nome:"))
    c += 1
print("O escolhido é {}".format(random.choice(nome)))
