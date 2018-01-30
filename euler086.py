from euler import timer, issquare
from sympy import isprime


# This takes >1000seconds, needs to be improved
@timer
def euler86():
    limit = 1000000
    n = 1
    total = 0

    def a(n):
        return sum(sum(1 for b in range(a, n + 1) if issquare((a + b)**2 + n**2)) for a in range(1, n + 1))
    
    while True:
        n += 1
        if n > 3 and isprime(n):
            continue
        current = a(n)
        total += current
        if total > limit:
            print(n)
            return


euler86()
