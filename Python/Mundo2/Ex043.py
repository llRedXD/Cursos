nome = input("Qual seu nome? ")
Kg = int(input("Qual sua massa(Kg): "))
A = float(input("Qual a sua altura: "))
IMC = Kg / (A**2)

print(f"Seu IMC é de {IMC:.2f}")
if IMC <= 18.4:
    print(f"{nome} esta abaixo do seu peso ideal!")
elif IMC >= 18.5 and IMC <= 24.9:
    print(f"{nome} esta no seu peso ideal")
elif IMC >= 25 and IMC <= 29.9:
    print(f"{nome} esta com sobrepeso!")
elif IMC >= 30 and IMC <= 34.9:
    print(f"{nome} esta no estagio I de obesidade!")
elif IMC >= 35 and IMC <= 39.9:
    print(f"{nome} esta no estagio II de obesidade!")
elif IMC >= 40:
    print(f"{nome} esta no estagio III de obesidade!")
