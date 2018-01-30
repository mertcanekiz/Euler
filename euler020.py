from euler import timer
from math import factorial


@timer
def euler20():
    print(sum(int(c) for c in str(factorial(100))))


euler20()
