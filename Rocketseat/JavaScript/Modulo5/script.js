// 1. Declare uma variavel de nome weight

let weight = 88.5

// 2. Que tipo de dado é a variavel acima?

console.log(typeof weight)

/* 
   3. Declare uma variavel e atribua valores para cada um dos dados:
    name: String
    age: Number(integer)
    stars: Number(float)
    isSubscribed: Boolean
*/

const film = {
    name: "Spider-Man",
    age: 2021,
    stars: 4.6,
    isSubscribed: true
}

/* 
   4. A variavel student abaixo é de que tipo de dados?

   4.1 Atribua a ela as mesmas propriedades e valores do exercicio 3

   4.2 Mostre no console a seguinte mensagem:

        <name> de idade <age> pesa <weight> kg.
*/

let student = {
    name: "Jorge",
    age: 17,
    stars: 4.6,
    isSubscribed: true,
    weight: 71
}

console.log(typeof student)
console.log(`${student.name} de idade ${student.age} pesa ${student.weight}` )

/* 
   5. Declare uma variavel do tipo array, de nome students e atribua nenhum valor, ou seja somente um array vazio.
*/
let students = []

/* 
   6. Reatribua valor para a variavel acima, colocando dentro dela o obejeto student da questão 4.
*/

students = [student]
console.log(students)

/* 
   7. Coloque no console o valor da posição 0 do array acima
*/

console.log(students[0])

/* 
   8. Crie um novo students e coloque na posição 1 do array students
*/

let student1 = {
    name: "Claudio",
    age: 26,
    stars: 4.6,
    isSubscribed: true,
    weight: 78
}

students[1] = student1
console.log(students)

/* 
   9. Sem rodar o codigo responda qual é a resposta do codigo abaixo e por que? Apos sua resposta, rode o codigo e veja se voce acertou
*/

// R = undefined, pois o JS faz um hoisting assim mesmo a variavel não tendo sido criada o JS leva ela para cima mas sem um valor atribuido 

console.log(a)
var a = 1