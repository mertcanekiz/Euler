from euler import timer, prime_partition
from itertools import count


@timer
def euler77():
    print(next(i for i in count(start=2) if prime_partition(i) > 5000))


euler77()
