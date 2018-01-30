from euler import timer, partition
from itertools import count


@timer
def euler78():
    print(next(i for i in count() if partition(i) % 1000000 == 0))


euler78()
