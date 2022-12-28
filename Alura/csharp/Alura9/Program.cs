using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Executando o Projeto 10");

        double investimento = 1000;

        int mes = 1;

        while (mes <= 12)
        {
            investimento = investimento + investimento * 0.005;
            Console.WriteLine("No mês " + mes + " você tem R$" + investimento);
            mes += 1;
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}