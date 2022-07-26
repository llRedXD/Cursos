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


def multiple_of(base, n): return n % base == 0


multiple_of_5 = partial(multiple_of, 5)
multiple_of_3 = partial(multiple_of, 3)


def assert_equal(result, expected):
    from sys import _getframe
    if result != expected:
        current = _getframe().f_back
        line_no = current.f_lineno
        msg = "Fail: Line {} got {} expecting {}"
        print(msg.format(line_no, result, expected))


