num1 = int(input("Digite um valor:"))
num2 = int(input("Digite um valor:"))

if num1 > num2:
    print(f"O primeiro numero é maior que o segundo")
elif num2 > num1:
    print(f"O segundo numero é maior que o primeiro")
elif num1 == num2:
    print("Os numeros são iguais")