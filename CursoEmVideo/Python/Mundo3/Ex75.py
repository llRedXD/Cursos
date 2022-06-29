num = []
cont = 0
par = 0
nove = 0
posi3 = 0

print("Investigador de Tuplas")
for i in range(1,5):
    dig = int(input("Digite um valor:"))
    num.append(dig)
c = tuple(num)
for i in num:
    div = i % 2
    print(f"Div:{div}")
    if div == 0:
        par = par + 1
    if i == 9:
        nove = nove + 1

posi3 = c.index(3)    
print(f"Noves:{nove}")
print(f"Posição 1º 3:{posi3}")
print(f"Par:{par}")

print(c)

