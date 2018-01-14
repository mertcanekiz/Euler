from euler import timer
from sympy import isprime, sieve


def prime_streak(a, b):
    n = 0
    while isprime(n * n + a * n + b):
        n += 1
    return n


@timer
def euler27():
    print(max(((prime_streak(a, b), a * b) 
               for b in sieve.primerange(2, 1000) for a in range(-b+2, 0, 2)))[1])


euler27()
