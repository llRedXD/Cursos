import random

numero_secreto = random.randrange(0,5)
numero = int(input("Qual o numero secreto:"))
if numero_secreto == numero:
    print("Parabens!")
else:
    print("Errou")
