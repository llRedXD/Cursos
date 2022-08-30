lista = []


def listaUnica():
    continua = True
    while continua:
        lista = criaLista()
        if input("Deseja continuar? (S/N) ").upper() == "N":
            continua = False
            lista.sort()
    print(f"o numeros da lista sao {lista or 'Lista vazia'}")

def criaLista():
    try:
        num = int(input("Digite um numero: "))
        if num not in lista:
            lista.append(num)
        else:
            print("Numero ja existe na lista")
    except Exception:
        print("Digite um numero valido")
    return lista
        
listaUnica()
