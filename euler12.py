from euler import timer, triangle
from sympy import sieve, divisor_count


@timer
def euler12():
    num_divisors = 0
    i = 0
    while num_divisors < 500:
        i += 1
        current = triangle(i)
        num_divisors = divisor_count(current)
    print(current, num_divisors)


euler12()
