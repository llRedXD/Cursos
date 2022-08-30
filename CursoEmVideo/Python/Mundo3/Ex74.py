import random


c = []

for _ in range(5):
    ram = (random.randrange(0,100))
    dom = c.append(ram)

c.sort(key = int)
num = tuple(c)
print("Maior Numero:",num[4])
print("Menor Numero:",num[0])

numb = tuple(random.randrange(0, 100) for _ in range(5))
print(numb)
sorted(numb)
print(f"Maior Numero: {numb[4]} Menor Numero: {numb[0]}")
