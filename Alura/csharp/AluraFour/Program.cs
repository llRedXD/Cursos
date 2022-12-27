using System;

class Program
{
    static void Main(string[] args)
    {
        char letra = 'b';
        Console.WriteLine("Letra: " + letra);

        letra = (char)65; // 65 é o código ASCII da letra A
        Console.WriteLine("Letra: " + letra);

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}