percurso = int(input("Qual o tamanho do percurso:"))
if percurso > 200:
    preço = percurso * 0.45
    print("O preço da passagem fica {}".format(preço))
else:
    preço = percurso * 0.50
    print("O preço da passagem fica {}".format(preço))