c = 0
primeiro = int(input("Digite o primeiro termo da PA:"))
razao = int(input("Digite a razão:"))
decimo = primeiro + (10)* razao
print(f"RAZÃO:", end="")
for c in range(primeiro,decimo,razao):
    print(f"{c}\033[31;40m->\033[m",end="")
print("FIM")