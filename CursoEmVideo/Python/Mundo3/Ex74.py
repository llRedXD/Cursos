import random


c = []

for i in range(0,5):
    ram = (random.randrange(0,100))
    dom = c.append(ram)


# print(c)
c.sort(key = int)
# print(c)
num = tuple(c)
print("Maior Numero:",(num[4]))
print("Menor Numero:",num[0])