# Ler 5 numero e depois armazenar em uma lista e mostrar o maior e o menor numero e suas respectivas posições

def main():
    num = [int(input(f"Digite o {i + 1}º numero: ")) for i in range(5)]
    print(f"O menor numero é {min(num)} e está na posição {num.index(min(num)) + 1}")
    print(f"O maior numero é {max(num)} e está na posição {num.index(max(num)) + 1}")

    
main()

