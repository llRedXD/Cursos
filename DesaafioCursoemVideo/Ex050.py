soma = 0
for c in range(0,7):
    c = int(input("Digite um valor:"))
    if c%2 == 0:
        soma = soma + c
print(f"A SOMA DOS VALORES:{soma}")