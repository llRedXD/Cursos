print("[M] Masculino\n[F] Feminino")
escolha = str(input("")).strip().upper()[0]
while escolha not in "MF":
    print("INFOMAÇÕES INCORRETAS\n[M] Masculino\n[F] Feminino")
    escolha = str(input()).strip().upper()[0]
if escolha == "M":
    print("Masculino")
elif escolha == "F":
    print("Feminino")
