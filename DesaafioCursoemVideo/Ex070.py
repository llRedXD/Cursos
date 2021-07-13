resp = ""
nome = ""
preço = 0
nomebar = ""
maismil = 0
preçom = 0
tot = 0

print("=-="*5)
print(" "*4,"CAIXA"," "*5)
print("=-="*5)
while resp != "N":
    print("\033[1;31m----PRODUTO----\033[0;0m")
    nome = str(input("Qual o nome do produto:"))
    preço = float(input("Qual o preço: R$"))
    if preço >= 1000:
        maismil = maismil + 1
    tot = tot + preço   
    resp = str(input("Você quer continuar [S]/[N]")).upper()
    if resp != "":
        preçom = preço
        if preço < preçom:
            preçom = preço
            nomebar = nome


print("\n") 
print("---"*5)   
print("   ORÇAMENTO   ")
print("---"*5)
print("\n")
print(f"Total a pagar: R${tot:.2f}")
print(f"Produtos com preço maior de R$1000,00: {maismil}Unidades")
print(f"O produto mais barato: {nomebar.capitalize()} Preço: R${preçom:.2f}")
print("\n")
