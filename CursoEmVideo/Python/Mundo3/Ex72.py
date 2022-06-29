Tupla = ("Zero","Um","Dois","Três","Quatro","Cinco","Seis","Sete","Oito","Nove","Dez","Onze","Doze","Treze","Quatorze","Quinze","Dezesseis","Dezesete","Dezoito","Dezenove","Vinte")
resp = ""
while  (1):
    numero = int(input("Qual numero você quer ver:"))
    if numero > 0 and numero < 20:
        print(Tupla[numero])
    
