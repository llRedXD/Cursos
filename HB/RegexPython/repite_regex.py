from re import search, match, findall

# Definir Quantidade
print(match(r'\d{4}', "12345"))

# Quantidade Minima
print(match(r'\d{2,}', "1234"))

# Repetição Preguiçosa
print(match(r'\d{2,}?', "1234"))

# Quantidade Minima e Maxima
print(match(r'\d{2,4}', "12345"))

# 0 ou 1 Ocorrencia
print(match(r'\d{,1}', "123456"))

# preguiçosa
print(match(r'\d??', "123456"))

# 0 ou mais ocorrencias
print(match(r'\d*', "12345"))

# 1 ou mais ocorrencias
print(match(r'\d+', "12345"))


# Exemplos para uso de repetições preguiçosas
text1 = 'attr1="value1" attr2="value2"'

print(findall(r'".+?"', text1))

text2 = 'attr1="" attr2=""'

print(findall(r'".*?"', text2))

# Match Object

m = match(r"\d+", "12345")
print(type(m))
print(m.group())
print(m.start())
print(m.end())
print(m.span())

# O Match Object so é retornado quando ele encotra algo senão retorna None
print(match(r"\d+", "abc"))

# Grupo de Captura

html1 = '<input type="text" id="id_cpf" name="John">'

html2 = '<input  id="cpf_id" type="txet" name="nhoJ">'

pattern = '<(.+?) type="(.+?)" id="(.+?)" name="(.+?)">'

m = match(pattern, html1)  # retorna o match object

# Traz todos os grupos de captura
print(m.groups())

# Traz um ou mais grupos especificos
print(m.group(1, 4))

pattern = '<(?.+?) (?:(?:type="(?.+?)"|id="(?.+?)"|name="(?.+?)") ?)*>'

m = match(pattern, html2)

print(m.groupdict())
