algoritmo "JOKENP O"
var
   jok:vetor[0..2 ] de caracter
   cv,jv,emp,comp ,esc:inteiro
   jogador,resp:caractere
inicio
   Repita
      jok[0] <- " PEDRA"
      jok[1] <- " TESOURA"
      jok[2] <- " PAPEL"
      comp <- randi(3)
      Repita
         Escreval ("----------JOKEN-PO!---------")
         Escreval ("-----------PLACAR-----------")
         Escreval ("JOGADOR     - ",jv)
         Escreval ("COMPUTADOR  - ",cv)
         Escreval ("EMPATE      - ",emp)
         Escreval ("----------- -----------------")
         Escreval ("QUAL A SUA  ARMA?")
         Escreval ("[1] PEDRA")
         Escreval ("[2] PAPEL")
         Escreval ("[3] TESOURA ")
         Leia(esc )
         Escolha( esc)
         Caso 1
            jogador <- "PEDRA"
         Caso 2
            Jogador <- "PAPEL"
         Caso 3
            Jogador <- "TESOUR A"
         Outrocaso
            Escreva("NÃO É UMA  OPÇÃO VALIDA !!")
         FimEscolha
         timer 500 on
         Limpatela
         timer 0
      Ate (esc=1)  ou (esc=2) o u (esc=3)
      timer 800 on
      Escreval("JO")
      Escreval("KEN")
      Escreval("P O!!!!")
      Limpatela
      EscrevaL("JOGADOR USOU " ,Jogador)
      EscrevaL("C OMPUTADOR USOU ",jok[comp] )
      timer 0 off
      Se (comp =  0) e (jogador  = "PEDRA") entao
         Escreval ("EMPATE!!")
         emp <- emp + 1
      Senao
         Se (comp  = 0) e (jogador = "TESOURA") entao
            Escreval("O COMPUTADOR VENCEU!! ")
            cv <-  cv + 1
         Senao
            Se (comp = 0) e (jogador = "PAPEL") entao
               Escreval("VOCÊ  VENCEU!!")
               jv  <- jv + 1
            FimSE
         FimSe
      FimSe
      Se (comp =  1) e (jogador  = "PEDRA") entao
         Escreval ("VOCÊ VENCEU !!")
         jv <- jv  + 1
      Senao
         Se (comp  = 1) e (jogador = "TESOURA") entao
            Escreval("EMPATE!! ")
            emp < - emp + 1
         Senao
            Se (comp = 1) e (jogador = "PAPEL") entao
               Escreval("O COMPUTADOR VENCEU!!")
               cv  <- cv + 1
            FimSE
         FimSe
      FimSe
      Se (comp =  2) e (jogador  = "PEDRA") entao
         Escreval ("O COMPUTADO R VENCEU!!")
         cv <- cv  + 1
      Senao
         Se (comp  = 2) e (jogador = "TESOUR A") entao
            Escreval("VOCÊ VENCEU!!")
            jv <-  jv + 1
         Senao
            Se (comp = 2) e (jogador = "PAPEL") entao
               Escreval("EMPATE!!")
               emp <- emp + 1
            FimSE
         FimSe
      FimSe
      escreva("Quer jogar de n ovo? S/N")
      Leia(resp)
      Limpatela
   Ate (resp = "N ")
fimalgoritmo