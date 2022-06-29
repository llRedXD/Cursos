from re import search, match, findall
import re

text ='\\section\n'
print(match('\\section', text))
print(match(r'\\section', text))

