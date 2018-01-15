from euler import timer
from sympy import sieve, isprime


def truncations(x):
    return [(x[:i], x[i:]) for i in range(len(x))]


def is_truncatable(x):
    return all(isprime(int(p)) and isprime(int(q)) for p, q in truncations(x) if p != '' and q != '')


@timer
def euler37():
    print(sum(p for p in sieve.primerange(10, 1000000) if is_truncatable(str(p))))


euler37()
