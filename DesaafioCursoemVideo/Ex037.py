print("-=-"*10)
print("CONVERSOR DE BASES NUMERICAS")
print("-=-"*10)
numero = int(input("Digite numero:"))
print("-=-"*10)
print("[1] Binario")
print("[2] Octal")
print("[3] Hexadecimal")
escolha = input()

if escolha == 1:
    print(f"O valor em base binaria é {bin(numero)}")
elif escolha == 2:
    print(f"O valor em base octal é {oct(numero)}")
elif escolha == 3:
    print(f"O valor em hexadecimal é {hex(numero)}")