from functools import wraps, lru_cache
from math import gcd
from time import time


def timer(func):
    @wraps(func)
    def wrap(*args, **kwargs):
        start = time()
        result = func(*args, **kwargs)
        end = time()
        print('%r took %2.4fs' % (func.__name__, end - start))
        return result
    return wrap


@lru_cache(maxsize=None)
def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fib(n - 1) + fib(n - 2)
# def fib(n):
# 	a, b = 0, 1
# 	for i in range(n):
# 		a, b = b, a+b
# 	return a


def largest_prime_factor(n):
    largest = 0
    counter = 2
    while counter * counter <= n:
        if n % counter == 0:
            n //= counter
            largest = counter
        else:
            counter += 1

    if n > largest:
        largest = n

    return largest


def lcm(a, b):
    return a // gcd(a, b) * b


def isprime(n):
    if n == 1:
        return False
    if n == 2 or n == 3:
        return True
    if n % 2 == 0:
        return False
    if n == 5 or n == 7:
        return True
    if n % 3 == 0:
        return False
    f = 5
    while f * f <= n:
        if n % f == 0:
            return False
        if n % (f + 2) == 0:
            return False
        f += 6
    return True


def product(iterable):
    result = 1
    for element in iterable:
        result *= int(element)
    return result


def triangle(n):
    return n * (n + 1) // 2


def choose(n, r):
    result = 1

    for i in range(1, r + 1):
        result *= (n - r + i)
        result //= i

    return result


def is_pandigital(n):
    n = str(n)
    s = len(n)
    return not '1234567890'[:s].strip(n)


def alphabetical_value(s):
    return sum(ord(c) - ord('A') + 1 for c in str(s))


def is_permutation(*args):
    args = iter(args)
    try:
        first = next(args)
    except StopIteration:
        return True
    return all(sorted(str(first)) == sorted(str(rest)) for rest in args)


def concat(a, b):
    return int(str(a)+str(b))
