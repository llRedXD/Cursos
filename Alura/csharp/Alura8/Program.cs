using System;

class Program
{
    static void Main(string[] args)
    {
        int gabrielIdade = 11;
        int acompanhante = 2;

        bool grupo = acompanhante > 0;

        string textoAdicional;

        if (grupo)
        {
            textoAdicional = "Gabriel está acompanhado";
        }
        else
        {
            textoAdicional = "Gabriel não está acompanhado";
        }


        if (gabrielIdade > 18 || grupo)
        {
            Console.WriteLine(textoAdicional);
            Console.WriteLine("ENTROU");
        }
        else
        {
            Console.WriteLine(textoAdicional);
            Console.WriteLine("NÃO ENTROU");
        }

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}