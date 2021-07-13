print("*"*30)
print("ALISATAMENTO")
print("*"*30)
idade = int(input("Qual sua idade:"))
if idade < 18:
    print("Você nao pode se alistar!")
elif idade >= 18:
    print("Você pode se alistar!")
elif idade > 45:
    print("Voce estorou o prazo para se alistar!")