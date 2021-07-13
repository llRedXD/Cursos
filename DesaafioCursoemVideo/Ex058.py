import random

Erros = 0
numero_secreto = random.randrange(0, 10)
numero = 0
acerto = 0


while acerto != 1:
    numero = int(input("Qual o numero secreto:"))
    if numero_secreto == numero:
        print("Parabens!")
        acerto += 1
    else:
        print("Errou")
        Erros += 1
print(f"Voce errou {Erros} vezes")