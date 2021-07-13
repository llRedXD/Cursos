frase = str(input("Digite uma frase:")).lstrip().upper()
print("A palavra 'A' aparece {} na frase".format(frase.count("A")))
print("A palavra 'a' aparece pela primeira vez na posição {}".format(frase.find("A")+1))
print("A palavra 'a' aparece pela ultima vez na posição {}".format(frase.rfind("A")+1))
