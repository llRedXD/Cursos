primeiro = int(input("Digite o primeiro termo da PA:"))
razao = int(input("Digite a razão:"))
decimo = primeiro + (10) * razao
print(f"RAZÃO:", end="")
while primeiro < decimo:
    print(f"{primeiro}\033[31;40m->\033[m", end="")
    primeiro += razao
print("")
decimo = 0
escolha = int(input("Voce quer ver mais quantos termos:"))
decimo += primeiro + escolha * razao
while primeiro < decimo:
    print(f"{primeiro}\033[31;40m->\033[m", end="")
    primeiro += razao
print("FIM")
