from euler import timer, is_permutation
from itertools import count


@timer
def euler52():
    print(next(i for i in count(start=99999, step=9) if is_permutation(i, 2*i, 3*i, 4*i, 5*i, 6*i)))


euler52()
