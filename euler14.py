from sys import setrecursionlimit
from euler import timer
from functools import lru_cache

def collatz2(n, acc=0):
    while True:
        if n < 2:
            return 1 + acc
        if n % 2 == 0:
            n, acc = n // 2, acc + 1
            continue
        else:
            n, acc = 3 * n + 1, acc + 1
            continue

@lru_cache(maxsize=None)
def collatz(n):
    if n == 1:
        return 1
    else:
        if n % 2 == 0:
            return 1 + collatz(n//2)
        else:
            return 1 + collatz(3*n+1)


setrecursionlimit(10000)
@timer
def euler14():
    maxchain, result = 0, 0
    limit = 1000000
    for i in range(limit, 0, -1):
        current = collatz(i)
        if current > maxchain:
            result = i
            maxchain = current
    print(result, maxchain)


euler14()
