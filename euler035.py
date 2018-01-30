from euler import timer
from sympy import sieve
from re import search


@timer
def euler35():
    primes = set(['2', '5'] + [str(p) for p in sieve.primerange(2, 1000000) if not search('[024568]', str(p))])
    print(sum(all(q in primes for q in [p[n:] + p[:n] for n in range(1, len(p))]) for p in primes))


euler35()
