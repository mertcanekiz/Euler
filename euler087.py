from euler import timer
from sympy import sieve


@timer
def euler087():
    result = set()
    limit = 50000000
    for p1 in sieve.primerange(2, limit**(1/2)):
        for p2 in sieve.primerange(2, limit**(1/3)):
            for p3 in sieve.primerange(2, limit**(1/4)):
                current = p1*p1 + p2*p2*p2 + p3*p3*p3*p3
                if current < limit:
                    result.add(current)
    print(len(result))


euler087()
