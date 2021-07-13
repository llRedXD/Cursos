algoritmo "semnome"
var
   Ano, N, Idade: inteiro
inicio
      escreval("------------------------")
      escreval("Departamento de Transito")
      escreval("------------------------")
      escreva("Em qual ano nasceu? (xxxx) ")
      leia(N)
      escreva("Em que anos estamos? ")
      leia(Ano)
      Idade <- Ano - N
            se(Idade >= 18)entao
            escreval("-----------------------------------")
            escreval("         Voce tem 18 anos          ")
            escreval("  Esta apto para Tirar a licença   ")
            escreval("-----------------------------------")
            senao
            escreval("--------------------------------------")
            escreval("       Voce nao tem 18 anos           ")
            escreval("  Nao sta apto para Tirar a licença   ")
            escreval("--------------------------------------")
            fimse
fimalgoritmo
