using System;

class Program
{
    static void Main(string[] args)
    {

        for (int contadorLinhas = 0; contadorLinhas <= 10; contadorLinhas++)
        {
            for (int contadorColuna = 0; contadorColuna < 10; contadorColuna++)
            {
                if (contadorColuna >= contadorLinhas)
                    break;
                Console.Write("*");
            }
            Console.WriteLine();
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}