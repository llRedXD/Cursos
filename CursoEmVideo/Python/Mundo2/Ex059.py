print("="*10)
print("CENTRAL")
print("="*10)
numero1 = int(input("Digite um valor:"))
numero2 = int(input("Digite um valor:"))

f = 0
while f < 1:
    print("="*10)
    print("   MENU")
    print("="*10)
    operacao = int(input("[1] SOMA\n[2] MULTIPLICAR\n[3] MAIOR\n[4] NOVOS NUMEROS\n[5] SAIR DO PROGRAMA\nEscolha:"))
    print("-=-"*3)
    if operacao == 1:
        soma = numero1 + numero2
        print(f"A soma é {soma}")
    elif operacao == 2:
        mult = numero1 * numero2
        print(f"A multiplicação é {mult}")
    elif operacao == 3:
        if numero1 > numero2:
            maior = numero1
            print(f"O maior numero é {maior}")
        elif numero1 < numero2:
            maior = numero2
            print(f"O maior numero é {maior}")
        elif numero1 == numero2:
            print("Sao iguais")
    elif operacao == 4:
        numero1 = int(input("Digite um valor:"))
        numero2 = int(input("Digite um valor:"))
    elif operacao == 5:
        print("Encerrando...")
        break
