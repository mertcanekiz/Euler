from euler import timer, is_permutation
from itertools import permutations
from sortedcontainers import SortedSet


# This has huge room for improvement; it runs at around 100s on my system
@timer
def euler62():
    cubes = SortedSet(i**3 for i in range(10000))
    for i1, c1 in enumerate(cubes):
        count = 1
        for c2 in cubes[i1+1:]:
            if is_permutation(c1, c2): count += 1
        if count == 5:
            print(c1)
            return


euler62()
