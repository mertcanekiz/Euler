from euler import timer
from math import gcd


@timer
def euler():
    fracs = set()
    for n in range(5, 12001):
        for k in range(n // 3 + 1, (n - 1) // 2 + 1):
            if gcd(n, k) == 1:
                fracs.add((n, k))
    print(len(fracs))


euler()