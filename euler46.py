from euler import timer
from sympy import sieve, isprime


def is_square(x):
    test = x**0.5
    return test == int(test)


@timer
def euler46():
    i = 7
    found = False
    while not found:
        found = True
        i += 2
        if isprime(i) or any(is_square((i - p) / 2) for p in sieve.primerange(2, i)):
            found = False
    print(i)


euler46()
