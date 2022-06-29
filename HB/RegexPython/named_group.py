from re import search, match, findall

html1 = '<input type="text" id="id_cpf" name="John">'

html2 = '<input  id="cpf_id" type="txet" name="nhoJ">'

pattern = '<(?P<tag>.+?) (?:(?:type="(?P<type>.+?)"|id="(?P<id>.+?)"|name="(?P<name>.+?)") ?)*>'

print(match(pattern, html1).groupdict())
