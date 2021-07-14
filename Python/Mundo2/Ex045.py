import random

print("\*/"*10)
print(" "*10,"JOKENPO")
print("\*/"*10)

bot = random.randrange(1,4)
if bot == 1:
    jogada_bot = "Pedra"
elif bot == 2:
    jogada_bot = "Tesoura"
elif bot == 3:
    jogada_bot = "Papel"

print("ESCOLHA SUA JOGADA")
print("[1] Pedra")
print("[2] Tesoura")
print("[3] Papel")
user = int(input())

if user == 1:
    jogada_user = "Pedra"
elif user == 2:
    jogada_user= "Tesoura"
elif user == 3:
    jogada_user = "Papel"

if bot ==  user:
    print(f"BOT:{jogada_bot} x USER:{jogada_user}\n     EMPATE!")
elif bot == 1 and user == 2 or bot == 2 and user == 1:
    print(f" BOT:{jogada_bot} x USER:{jogada_user}\n    DERROTA!")
elif bot == 1 and user == 3 or bot == 3 and user == 1:
    print(f"BOT:{jogada_bot} x USER:{jogada_user}\n     VITORIA!")
elif bot == 2 and user == 3 or bot == 3 and user == 2:
    print(f"BOT:{jogada_bot} x USER:{jogada_user}\n     VITORIA!")
