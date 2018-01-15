from euler import timer
from math import gcd


@timer
def euler39():
    max_solutions, max_p = 0, 0
    for p in range(2, 1001, 2):
        solutions = 0
        for a in range(2, p // 3):
            for b in range(a, p // 2):
                c = p - b - a
                if a * a + b * b == c * c:
                    solutions += 1
        if solutions > max_solutions:
            max_solutions, max_p = solutions, p
    print(f"{max_p} yields {max_solutions} solutions")


euler39()
