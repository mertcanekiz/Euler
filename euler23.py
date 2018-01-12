from euler import timer
from sympy import divisors


@timer
def euler23():
    limit = 28123
    abundant_numbers = [i for i in range(limit) if sum(divisors(i)) - i > i]
    abundant_sums = [False] * (limit + 1)
    for i, a1 in enumerate(abundant_numbers):
        for j, a2 in enumerate(abundant_numbers[i:]):
            current = a1 + a2
            if current <= limit:
                abundant_sums[current] = True
            else:
                break
    print(sum(i for i, s in enumerate(abundant_sums) if not s))


euler23()
