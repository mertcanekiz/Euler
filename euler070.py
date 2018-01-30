from euler import timer, is_permutation
from sympy import sieve


@timer
def euler70():
    result = 1
    min_ratio = 2
    limit = 10000000
    lowerbound, upperbound = int(0.7 * (limit)**0.5), int(1.3 * (limit)**0.5)
    for p in sieve.primerange(lowerbound, upperbound):
        for q in sieve.primerange(p + 1, upperbound):
            n = p * q
            if n > limit:
                break
            phi = (p - 1) * (q - 1)
            ratio = n / phi
            if is_permutation(n, int(phi)) and min_ratio > ratio:
                result, min_ratio = n, ratio
    print(result)


euler70()
