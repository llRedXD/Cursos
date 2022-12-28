using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Executando o Projeto 11");

        double investimento = 1000;
        double taxaDeRentabilidade = 1.005;
        

        for (int i = 1; i <= 5; i++)
        {
            Console.WriteLine("Ano " + i);
            
            for (int r = 1; r <= 12; r++)
            {
                investimento *= taxaDeRentabilidade; // investimento = investimento + investimento * 0.0005
                Console.WriteLine("No mês " + r + " você tem R$" + investimento);
            }
            taxaDeRentabilidade += 0.001;
        }

        

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}