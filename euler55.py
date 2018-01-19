from euler import timer
from functools import lru_cache


def lychrel(n):
    for i in range(50):
        n = n + int(str(n)[::-1])
        if int(str(n)[::-1]) == n:
            return False
    return True


@timer
def euler55():
    print(sum(1 for x in range(1, 10000) if lychrel(x)))


euler55()
