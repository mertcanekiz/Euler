from euler import timer, is_pandigital
from sympy import sieve


@timer
def euler41():
    print(max(p for p in sieve.primerange(2, 7654321) if is_pandigital(p)))


euler41()
