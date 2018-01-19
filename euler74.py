from euler import timer
from math import factorial


def digit_factorial(n):
    total = 0
    while n > 0:
        total += factorial(n % 10)
        n //= 10
    return total


def chain_length(n):
    sequence = []
    while n not in sequence:
        sequence.append(n)
        n = digit_factorial(n)
    return len(sequence)


@timer
def euler():
    print(sum(1 for i in range(1, 1000001) if chain_length(i) == 60))


euler()