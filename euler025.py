from euler import timer, fib
from itertools import count


@timer
def euler25():
    print(next(i for i in count() if len(str(fib(i))) >= 1000))


euler25()
