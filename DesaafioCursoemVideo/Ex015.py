print("="*18)
print("Aluguel de carros")
print("="*18)
d = int(input("Quantos dias voce alugou o carro:"))
km = int(input("Quantos KM voce andou:"))
preço = (60 * d) + (km * 0.15)
print("Voce tem q paga R${:.2f}".format(preço))