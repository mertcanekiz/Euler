# Project Euler Problem 1
# https://projecteuler.net/problem=1

from euler import run

def euler():
    result = 0
    for i in range(1, 1000):
        if i % 3 == 0 or i % 5 == 0:
            result += i
    print result

run(euler)
