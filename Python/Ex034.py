sal = int(input("Qual o seu salario:"))
if sal > 1250:  
    novosal = sal + (sal*0.10)
else:
    novosal = sal + (sal*0.15)
print("Seu novo salario é de {}".format(novosal))