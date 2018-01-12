from euler import timer
from sympy import divisors


@timer
def euler21():
    divisor_sums = []
    limit = 10000
    results = set()
    for i in range(limit):
        divisor_sums.append(sum(divisors(i)) - i)
    for i, d in enumerate(divisor_sums):
        if d > len(divisor_sums):
            continue
        if divisor_sums[d] == i and i != d:
            results.add(i)
    print(sum(results))


euler21()
