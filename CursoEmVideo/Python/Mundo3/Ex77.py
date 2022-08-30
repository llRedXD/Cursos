palavras = {"Cinza", "Ter", "Cassino", "Soco",
            "Conectar", "Pegada", "Concha", "Verruga", "Ato", "Fogao"}

vogal = []
for palavra in palavras:
    vogal.extend(letra for letra in palavra if letra in "aeiou")
    print(f"Na palavra {palavra:<5} as vogais são ", end="")
    for v in vogal:
        print(f"{v}", end=" ")
    print("\n")
    vogal.clear()
    