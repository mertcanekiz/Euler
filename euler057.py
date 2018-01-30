from euler import timer
from math import floor, log10


def convergent(n):
    numerator = 1
    denominator = 2
    for i in range(n - 1):
        numerator += denominator * 2
        numerator, denominator = denominator, numerator
    numerator += denominator
    return numerator, denominator


@timer
def euler57():
    print(sum(1 for i in range(1, 1001) if len(str(convergent(i)[0])) > len(str(convergent(i)[1]))))


euler57()
