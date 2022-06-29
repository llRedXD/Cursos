from re import search, match, findall
import re

# Retorna Todos os itens da lista que estejam  no intervalo de caracteres
print(findall('[aeiou]', 'Gabriel Neira'))

# Retorna Todos os caracteres menos os listados
print(findall('[^aeiou]', 'Gabriel Neira'))

# Retorna Todos os caracteres que estejam no intervalo de caracteres 
print(findall('[a-z]', 'Gabriel Neira'))

# Sequencias especiais
    # \d == [0-9]
    # \D == [^0-9]
    # \s == [ \t\n\r\f\v]
    # \S == [^ \t\n\r\f\v]
    # \w == [a-zA-Z0-9_]
    # \W == [^a-zA-Z0-9_]
print(findall('\w', 'Gabriel Neira'))


