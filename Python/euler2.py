# Project Euler Problem 2
# https://projecteuler.net/problem=2

from euler import run, fib

def euler():
    limit = 4e6
    result = 0
    current = 0
    i = 0
    while current < limit:
        i += 1
        current = fib(i)
        if current % 2 == 0:
            result += current
    print(result)

run(euler)
