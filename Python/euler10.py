from sympy import sieve
from euler import timer


@timer
def euler10():
    print(sum(sieve.primerange(2, 2e6)))


euler10()
