from euler import timer, is_permutation
from sympy import sieve, isprime


@timer
def euler49():
    for p in sieve.primerange(1488, 9999):
        for i in range(1, 9999-p):
            if is_permutation(p, p+i) and is_permutation(p, p+2*i) and isprime(p+i) and isprime(p+2*i):
                print(str(p)+str(p+i)+str(p+2*i))
                return


euler49()
