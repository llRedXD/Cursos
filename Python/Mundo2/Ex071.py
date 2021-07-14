cinquenta = 0
vinte = 0
dez = 0
um = 0

print("=-="*6)
print(" CAIXA ELETRONICO")
print("=-="*6)
valor = int(input("Qual a sacar: R$"))
if valor >= 50:
    cinquenta = valor//50
    valor = valor % 50
    if valor >= 20:
        vinte = valor//20
        valor = valor%20
        if valor >= 10:
            dez = valor//10
            valor = valor%10
            if valor >= 1:
                um = valor//1

print("---"*5)
print("SAQUE")
print(f"NOTAS DE R$50: {cinquenta}")
print(f"NOTAS DE R$20: {vinte}")
print(f"NOTAS DE R$10: {dez}")
print(f"NOTAS DE R$1:  {um}")