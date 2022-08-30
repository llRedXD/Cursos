tupla = {"Sorvete":10.00, "Pão":2.00, "Café":5.00, "Cachorro Quente":4.00, "Refrigerante":3.00}

print("-"*30)
print(" "*6 + "Listagem de preços"+" " *6)
print("-"*30)
for i in tupla:
    print(f"{i:<17}R$ {tupla[i]:>5.2f}")