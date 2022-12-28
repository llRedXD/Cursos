using System;
using System.IO;

class Program
{
    static void Main(string[] args)
    {
        // Criar uma variavel para pegar o valor inserido pelo usuario
        string nome;

        // Mostrar uma mensagem para o usuario
        Console.WriteLine("Digite seu nome: ");

        // Pegar o valor inserido pelo usuario
        nome = Console.ReadLine();

        switch (nome) {
            case "João":
                Console.WriteLine("João 1");
                break;
            case "Gabriel":
                Console.WriteLine("Gabriel 1");
                break;
            default:
                Console.WriteLine("Nenhum Nome Escolhido");
                break;

        }

        // Mostrar o valor inserido pelo usuario
        Console.WriteLine("Seu nome é: " + nome);

        Console.WriteLine("Aperte enter para sair...");
        Console.ReadLine();
    }
}