print("="*20)
print("CASAS BAHIA")
print("="*20)
print("CAIXA")
preço_normal = float(input("Qual o preço:"))
print("[1]DINHEIRO/CHEQUE")
print("[2]CARTÃO")
escolha = int(input())
if escolha == 1:
    preço_din = preço_normal - (preço_normal * 0.1)
    print(f"Preço com desconto de 10% R${preço_din}")
elif escolha == 2:
    print("[1] A VISTA 5% DE DESCONTO")
    print("[2] 2x PREÇO NORMAL")
    print("[3] 3x OU MAIS 20% DE JUROS")
    sub_escolha = int(input())
    if sub_escolha == 1:
        preço_cartao = preço_normal - (preço_normal * 0.05)
        print(f"Preço R${preço_cartao}")
    elif sub_escolha == 2:
        preço_cartao = preço_normal
        print(f"Preço R${preço_cartao}")
    elif sub_escolha == 3:
        preço_cartao = preço_normal + (preço_normal * 0.2)
        print(f"Preço R${preço_cartao}")