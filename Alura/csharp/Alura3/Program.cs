using System;

class Programa
{
    static void Main(string[] args)
    {
        double salario;
        salario = 3000.12; // Se declaram como inteiro não ira compilar
        Console.WriteLine("Salário: " + salario);

        double idade;
        idade = 7.0 / 2; // Se não colocar o .0, o resultado será 3 (inteiro) e não 3.5 (double)
        Console.WriteLine(idade);

        // Tipo de variáveis 
        // int - inteiro
        // double - decimal
        // float - decimal com menos precisão
        // short - inteiro com menos espaço
        // long - inteiro com mais espaço
    }
}