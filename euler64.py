from euler import timer
from math import floor


def continued_fraction(n):
    if int(n**0.5) == n**0.5:
        return {n**0.5: []}
    frac = {}
    m, d, a0 = 0, 1, floor(n**0.5)
    a = a0
    frac[a0] = []
    while a != 2 * a0:
        m = d * a - m
        d = (n - m**2) // d
        a = (a0 + m) // d
        frac[a0].append(a)
    return frac


@timer
def euler64():
    print(sum(1 for n in range(2, 10001) if len(continued_fraction(n)[floor(n**0.5)]) % 2 == 1))


euler64()
