// Funcões construtoras

function Person(name) {
    this.name = name
    this.walk = () => {
        return "Andando"
    }
}

const santos = new Person("Santos")
const rafa = new Person("Rafaela")
console.log(santos.name + " está ",santos.walk())
console.log(rafa.name + " não esta ",rafa.walk())
