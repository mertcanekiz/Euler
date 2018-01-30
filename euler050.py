from euler import timer
from sympy import sieve, isprime
from sortedcontainers import SortedSet


@timer
def euler50():
    limit = 100
    primes = SortedSet(sieve.primerange(1, limit))
    prime_sums = SortedSet()
    prime_sums.add(0)
    consecutive = 0
    for p in primes:
        prime_sums.add(prime_sums[-1] + p)
    for i in range(consecutive, len(prime_sums)):
        for j in range(i - consecutive - 1, -1, -1):
            if prime_sums[i] - prime_sums[j] > limit:
                break
            if isprime(prime_sums[i] - prime_sums[j]):
                consecutive = i - j
                result = prime_sums[i] - prime_sums[j]
    print(result)


euler50()
