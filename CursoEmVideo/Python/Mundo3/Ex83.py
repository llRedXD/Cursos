# Exercício Python 083: Crie um programa onde o usuário digite uma expressão qualquer que use parênteses. 
# Seu aplicativo deverá analisar se a expressão passada está com os parênteses abertos e fechados na ordem correta.
# Ex : ((a+b)+c) está ok

expressao = str(input('Digite uma expressão: '))
contador: int = 0

for char in expressao:
    if char == "(" or char == ")":
        contador += 1

print(contador)
if contador % 2 == 0:
    print('Expressão válida!')
else:
    print('Expressão inválida!')