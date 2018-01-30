from euler import timer
from sortedcontainers import SortedSet
from math import gcd, floor, ceil
from fractions import Fraction


@timer
def euler():
    fracs = SortedSet()
    limit = 10**6
    for d in range(5, limit+1):
        target = d*3/7
        for n in range(int(floor(target)), int(ceil(target))+1):
            if gcd(n, d) == 1:
                fracs.add(Fraction(n, d))
    result = 0
    compare = Fraction(3, 7)
    for i, f in enumerate(fracs):
        if f == compare:
            result = i-1
            break
    print(fracs[result])


euler()
