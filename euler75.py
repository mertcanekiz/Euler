from euler import timer, gen_pythagorean_triples
from collections import defaultdict


@timer
def euler75():
    limit = 1500000
    result = 0
    triangles = [0] * (limit + 1)
    for l in range(12, limit + 1):
        for a, b, c in gen_pythagorean_triples(l):
            p = a + b + c
            while p <= limit:
                triangles[p] += 1
                if triangles[p] == 1:
                    result += 1
                if triangles[p] == 2:
                    result -= 1
                p += a + b + c
    print(result)


euler75()
