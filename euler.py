from functools import wraps, lru_cache
from math import gcd, floor
from time import time
from sympy import factorint


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
    return int(str(a) + str(b))


def continued_fraction(n):
    if int(n**0.5) == n**0.5:
        return [n]
    m, d, a0 = 0, 1, floor(n**0.5)
    a = a0
    frac = [a0, []]
    while a != 2 * a0:
        m = d * a - m
        d = (n - m**2) // d
        a = (a0 + m) // d
        frac[1].append(a)
    return frac


def convergent(n, frac):
    a0 = 1
    b0 = 0
    a = frac[0]
    b = 1
    for i in range(n):
        bn = frac[1][i % len(frac[1])]
        a, a0 = bn * a + a0, a
        b, b0 = bn * b + b0, b
    return a, b


def totient(n):
    if isprime(n):
        return n - 1
    result = n
    for p in factorint(n).keys():
        result *= (1 - 1 / p)
    return result


def gen_pythagorean_triples(s):
    mlimit = int((s // 2)**0.5)
    for m in range(2, mlimit + 1):
        if s // 2 % m == 0:
            if m % 2 == 0:
                k = m + 1
            else:
                k = m + 2
            while k < 2 * m and k <= s // (2 * m):
                if s // (2 * m) % k == 0 and gcd(k, m) == 1:
                    d = 1  # s // 2 // (k * m)
                    n = k - m
                    a = d * (m * m - n * n)
                    b = 2 * d * m * n
                    c = d * (m * m + n * n)
                    if a + b + c == s:
                        yield a, b, c
                k += 2


def pentagonal(n):
    return n * (3 * n - 1) // 2


@lru_cache(maxsize=None)
def generalized_pentagonal(n):
    if n == 0:
        return pentagonal(0)
    if n % 2 == 0:
        return pentagonal(-n // 2)
    else:
        return pentagonal(n // 2 + 1)


@lru_cache(maxsize=None)
def partition(n):
    if n < 0:
        return 0
    if n == 0:
        return 1
    result = 0
    m = 0
    penta = 0
    while penta <= n:
        m += 1
        sign = -1 if (m - 1) % 4 > 1 else 1
        penta = generalized_pentagonal(m)
        current = sign * partition(n - penta)
        result += current
        if current == 0:
            break
    return result


def sopf(x):
    """
    Sum of prime factors of x
    """
    return sum(factorint(x).keys())


@lru_cache(maxsize=None)
def prime_partition(x):
    """
    Number of summations of primes that add up to x
    """
    if x == 1:
        return 0
    return (sopf(x) + sum(sopf(j) * prime_partition(x - j) for j in range(1, x))) // x


def issquare(n):
    return int(n**0.5) == n**0.5
