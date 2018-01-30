from euler import timer, concat
from sympy import sieve, isprime


def find_concatenable_primes(p, primes):
    for q in primes:
        if isprime(concat(p, q)) and isprime(concat(q, p)):
            yield q


@timer
def euler60():
    limit = 10000
    primes = set(sieve.primerange(2, limit))
    for p0 in primes:
        s0 = set(find_concatenable_primes(p0, primes))
        for p1 in s0:
            s1 = set(find_concatenable_primes(p1, primes))
            i1 = s0 & s1
            for p2 in i1:
                s2 = set(find_concatenable_primes(p2, primes))
                i2 = i1 & s2
                for p3 in i2:
                    s3 = set(find_concatenable_primes(p3, primes))
                    i3 = i2 & s3
                    if len(i3) == 1:
                        print(p0+p1+p2+p3+sum(i3))
                        return



euler60()
