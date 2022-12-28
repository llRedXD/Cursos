using System;
using System.IO;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Fatorial Requerido: ");

        int x, r = 1;
        x = Convert.ToInt32(Console.ReadLine());



        for (int i = 1; i <= x; i++)
        {
            r *= i;
        }
        Console.WriteLine("O fatorial de " + x + " é " + r);

    }
}