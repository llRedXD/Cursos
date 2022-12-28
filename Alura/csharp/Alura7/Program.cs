using System;

class Program
{
    static void Main(string[] args)
    {
        int gabrielIdade = 11;
        int acompanhante = 2;

        bool grupo = acompanhante > 0;
        Console.WriteLine(grupo);

        if (gabrielIdade > 18 || grupo)
        {
            Console.WriteLine("ENTROU");
        }
        else
        {
            Console.WriteLine("NÃO ENTROU");
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}