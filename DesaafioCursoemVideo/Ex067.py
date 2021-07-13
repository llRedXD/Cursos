c = 0
number = 0
while True:
    number = int(input("\033[1;31mQuer ver a tabuada de qual valor:"))
    if number < 0:
        break
    print("=====TABUADA======")
    for c in range(0, 11):
        print(f"\033[1;97m{number} x {c} = {number*c}")
        c += 1
    print("---"*4)
print("Encerrando programa...")