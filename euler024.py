from euler import timer
from itertools import permutations, islice


@timer
def euler24():
    print(''.join(c for c in next(islice(permutations('0123456789'), 999999, 1000000))))


euler24()
