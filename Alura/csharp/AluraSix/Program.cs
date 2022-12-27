using System;

class Program
{
    static void Main(string[] args)
    {
        int gabrielIdade = 11;
        int acompanhante = 2;

        if (gabrielIdade > 18)
        {
            Console.WriteLine("ENTROU");
        }
        else
        {
            if (acompanhante > 0)
            {
                Console.WriteLine("ENTROU COM ACOMPANHANTE");
            }
            else
            {
                Console.WriteLine("NÃO ENTROU");
            }
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}