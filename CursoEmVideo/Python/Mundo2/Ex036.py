print("-=-"*10)
valor_casa = int(input("Qual o valor da casa:"))
salario = int(input("Qual seu salario:"))
ano = int(input("Quantos anos vc vai pagar:"))
print("-=-"*10)
parcela = valor_casa / (ano*12)
if parcela <= (salario * 0.3):
    print(f'O valor de R${parcela:.2f}')
else:
    print("Você nao pode pegar o emprestimo")
