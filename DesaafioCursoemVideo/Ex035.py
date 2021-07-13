A = float(input("Digite o valor A:"))
B = float(input("Digite o valor B:"))
C = float(input("Digite o valor C:"))
if A < B + C and B < A + C and C < B + A:
    print("Essas medidas podem formar um triangulo")
else:
    print("Essas medidas nao podem formar um triangulo")
