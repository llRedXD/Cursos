ano = int(input("Digite um ano:"))
if ano%4 == 0:
    print("O ano de {} é bissexto".format(ano))
else:
    print("O ano de {} nao é bissexto".format(ano))