print("=-="*10)
print("BAIXA DAS NOTAS")
print("=-="*10)
nota1 = float(input("Digite a nota:"))
nota2 = float(input("Digite a nota:"))
print("=-="*10)
media = (nota1 + nota2) / 2
if media < 5.0:
    print("REPROVADO")
elif media >= 5.0 and media <= 6.9:
    print("RECUPERAÇÃO")
elif media >= 7:
    print("APROVADO")