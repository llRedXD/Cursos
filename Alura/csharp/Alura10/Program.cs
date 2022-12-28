using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Executando o Projeto 10");

        double investimento = 1000;


        for (int i = 1; i <= 12; i++)
        {
            investimento *= 1.005; // investimento = investimento + investimento * 0.0005
            Console.WriteLine("No mês " + i + " você tem R$" + investimento);
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}