from euler import timer
from math import log10


@timer
def euler63():
    print(sum(int(1//(1-log10(n))) for n in range(1, 10)))


euler63()
