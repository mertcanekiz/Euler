from euler import timer
from sympy import isprime


def corners(n):
    return ((2*n+1)**2, (2*n+1)**2 - 2*n, (2*n+1)**2 - 4*n, (2*n+1)**2 - 6*n)


@timer
def euler58():
    i = 0
    prime_count = 0
    total = 1
    while True:
        i += 1
        prime_count += sum(1 for c in corners(i) if isprime(c))
        total += 4
        if prime_count / total < 0.1:
            print(2*i+1)
            break


euler58()
