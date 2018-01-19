from euler import timer
from sympy import prime


@timer
def euler69():
    result = 1
    i = 1
    while result * prime(i) < 1000000:
        result *= prime(i)
        i += 1
    print(result)


euler69()
