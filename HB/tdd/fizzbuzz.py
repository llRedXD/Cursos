"""
se for múltiplo de 3, imprima fizz
se for múltiplo de 5, imprima buzz
se for múltiplo de 3 e 5, imprima fizzbuzz
se não for, imprima o número
"""

from functools import partial


def robot(n):
    say = str(n)
    if multiple_of_3(n) and multiple_of_5(n):
        say = "fizzbuzz"
    elif multiple_of_5(n):
        say = "buzz"
    elif multiple_of_3(n):
        say = 'fizz'
    return say


multiple_of = lambda base, n: n % base == 0
multiple_of_5 = partial(multiple_of, 5)
multiple_of_3 = partial(multiple_of, 3)



if __name__ == '__main__':
    assert robot(1) == '1'
    assert robot(2) == '2'
    assert robot(4) == '4'

    assert robot(3) == 'fizz'
    assert robot(6) == 'fizz'
    assert robot(9) == 'fizz'

    assert robot(5) == 'buzz'
    assert robot(10) == 'buzz'
    assert robot(20) == 'buzz'

    assert robot(15) == 'fizzbuzz'
    assert robot(30) == 'fizzbuzz'
    assert robot(45) == 'fizzbuzz'
