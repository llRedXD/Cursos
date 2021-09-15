Times = ("Atlético - MG",
         "Palmeiras",
         "Fortaleza",
         "Bragantino",
         "Flamengo",
         "Corinthians",
         "Atlético - GO",
         "Ceará",
         "Athletico - PR",
         "Internacional",
         "Santos",
         "São Paulo",
         "Juventude",
         "Cuiabá",
         "Bahia",
         "Fluminense",
         "Grêmio",
         "Sport",
         "América - MG",
         "Chapecoense")

c = 0

print("------G4------")
for i in range(0,4):
    c += 1
    print(f"{c} {Times[i]}")
    
print("\n")

c = 16

print("------Z4------")
for i in range(16,20):
    c += 1
    print(f"{c} - {Times[i]}")

print("\n")

print("-----Times-----")
G = sorted(Times)
for i in range(0,19):
    print(f"{G[i]}")


print("----Pesquisa----")
time = input("Qual o time desejado:")
if time in Times:
    localizar = Times.index(time)
    lugar = c = localizar + 1
    print(f"{lugar} - {Times[localizar]}")
