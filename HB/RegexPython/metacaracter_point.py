from re import search, match, findall
import re

# Metacaracter Point(.)

print(match(".", "abc"))

print(search(".","abc"))

print(findall(".","abc")) 

# Âncoras de regex ^ $

print(findall('^.', 'abc\ndef\nghi', re.MULTILINE))

print(findall('.$', 'abc\ndef\nghi', re.MULTILINE))

print(findall('a', 'abca'))

print(findall('^$', ''))
print(findall('^$', '\n', re.MULTILINE))
