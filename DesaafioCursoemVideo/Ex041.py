print("="*25)
print("CONFEDERAÇÃO DE NATAÇÃO")
print("="*25)
idade = int (input("Qual a sua idade:"))
if idade <= 9:
    print("CATEGORIA:Mirim")
elif idade <= 14:
    print("CATEGORIA:Infantil")
elif idade <= 19:
    print("CATEGORIA:Junior")
elif idade <= 20:
    print("CATEGORIA:Senior")
elif idade >= 21:
    print("CATEGORIA:Master")
