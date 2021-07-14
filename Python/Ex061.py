primeiro = int(input("Digite o primeiro termo da PA:"))
razao = int(input("Digite a razão:"))
decimo = primeiro + (10) * razao
print(f"RAZÃO:", end="")
while primeiro < decimo:
    print(f"{primeiro}\033[31;40m->\033[m", end="")
    primeiro += razao
print("FIM")
