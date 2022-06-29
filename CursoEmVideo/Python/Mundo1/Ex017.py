import math

print("="*15)
print("HIPOTENUSA")
print("="*15)
b = float(input("Digite o tamanho do cateto adjacente:"))
c = float(input("Digite o tamanho do cateto oposto:"))
s = b**2 + c**2
a = math.sqrt(s)
print("A hipotenusa é {:.1f}cm".format(a))