algoritmo "Fatorial"
var
   n, f, m: inteiro
inicio
      escreva("Digite um numero:")
      leia(n)
      f <- n
      m <- 1
          repita
          m <- m * f
          f <- f - 1
          ate(f < 1)
      escreva("O fatorial de ", n, " é ", m)
fimalgoritmo