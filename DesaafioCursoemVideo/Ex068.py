import random

numeroCPU =random.randrange(0,11)

escolha = ""
soma = 0
resposta = ""
vitorias = 0
print("JOGO PAR OU IMPAR")
print("=-="*6)
while True:
    escolha = str(input("PAR ou IMPAR? [P]/[I]  ")).upper()
    print("-"*30)
    number = int(input("Digite um valor:"))
    print("-"*60)
    soma = number + numeroCPU
    ParOuImpar = soma%2
    if ParOuImpar == 0:
        resposta = "PAR"
    else:
        resposta = "IMPAR"
    print(f"Você jogou {number} o computador jogou {numeroCPU}. Total de {soma} deu {resposta}")
    if escolha == "P" and resposta == "PAR":
        vitorias += 1
    elif escolha == "I" and resposta == "IMPAR":
        vitorias += 1
    else:
        break
print("="*60)
print(f"Voce perdeu. Ganhou {vitorias} vezes")