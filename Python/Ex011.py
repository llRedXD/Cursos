l = float(input("Qual a largura:"))
h = float(input("Qual a altura:"))
a = l * h
balde = 2
uso = a / balde
print("Voce vai usar {:.2f} baldes para pintar uma parede de {:.2f}M²".format(uso,a))