c = 0
number = int(input("Quantos termos da sequencia quer ver:"))
n1 = 0
n2 = 1
print(n1,n2,end=" ")
while c < number:
    n3 = n1 + n2
    print(n3,end=" ")
    n1 = n2
    n2 = n3
    c += 1
