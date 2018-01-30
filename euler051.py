from euler import timer
from sympy import sieve, isprime


def replace_count(p, r):
    return sum(1 for d in '0123456789' if int(str.replace(p, r, d)) > 100000 and isprime(int(str.replace(p, r, d))))


@timer
def euler51():
    for p in sieve.primerange(100000, 999999):
        s = str(p)
        if s.count('0') == 3 and replace_count(s, '0') == 8 \
            or s.count('1') == 3 and replace_count(s, '1') == 8 \
                or s.count('2') == 3 and replace_count(s, '2') == 8:
            print(s)
            break


euler51()
