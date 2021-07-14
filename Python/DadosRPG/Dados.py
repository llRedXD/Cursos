import random
from flask import Flask

app = Flask(__name__)

lista = []

resp = ""
quantidade = 0
escolha = 0
dado = 0

def D100():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D100 :",end=" ")
    for c in range(0,quantidade):
        dado = random.randrange(1, 101,)
        lista.append(dado)
    Formatação()
    lista.clear()

def D20():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D20 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 21)
        lista.append(dado)
    Formatação()
    lista.clear()

def D12():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D12 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 13)
        lista.append(dado)
    Formatação()
    lista.clear()

def D10():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D10 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 11)
        lista.append(dado)
    Formatação()
    lista.clear()

def D8():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D8 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 9)
        lista.append(dado)
    Formatação()
    lista.clear()

def D6():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D6 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 7)
        lista.append(dado)
    Formatação()
    lista.clear()

def D4():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D4 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 5)
        lista.append(dado)
    Formatação()
    lista.clear()

def D2():
    quantidade = int(input("Quantos dados:"))
    print(f"{quantidade}D2 :", end=" ")
    for c in range(0, quantidade):
        dado = random.randrange(1, 3)
        lista.append(dado)
    Formatação()
    lista.clear()

def Formatação():
    lista.sort()
    lista.reverse()
    for number in lista:
        print(f"{number}", end=" ")
    print(f"=>",sum(lista))

while resp != "N":
    print("<===========>")
    print("<===DADOS===>")
    print("<===========>")
    print("|D100[1] D20[2]|")
    print("|D12 [3] D10[4]|")
    print("|D8  [5] D6 [6]|")
    print("|D4  [7] D2 [8]|")

    escolha = int(input("Qual dado:"))
    if escolha > 8:
        print("OPÇÃO INVALIDA.TENTE NOVAMENTE")
    while escolha > 8:
        escolha = int(input("Qual dado:"))
        if escolha > 8: 
            print("OPÇÃO INVALIDA.TENTE NOVAMENTE")
    
    if escolha == 1:
        D100()
   
    if escolha == 2:
        D20()
    
    if escolha == 3:
        D12()
    
    if escolha == 4:
        D10()
    
    if escolha == 5:
        D8()
    
    if escolha == 6:
        D6()
    
    if escolha == 7:
        D4()
    
    if escolha == 8:
        D2()

    resp = str(input("Quer rodar denovo[S]/[N]: ")).upper()
    if resp != "S" and resp != "N":
        print("OPÇÃO INVALIDA.TENTE NOVAMENTE")
    while resp != "S" and resp != "N":
        resp = str(input("Quer rodar denovo[S]/[N]: ")).upper()
        if resp != "S" and resp != "N":
            print("OPÇÃO INVALIDA.TENTE NOVAMENTE")
print("Encerrando...")

