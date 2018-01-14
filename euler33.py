from euler import timer, gcd
from decimal import Decimal


@timer
def euler33():
    nominators = 1
    denominators = 1
    for a in range(10, 100):
        for b in range(a + 1, 100):
            if a % 10 == 0 and b % 10 == 0:
                continue
            value = Decimal(a) / Decimal(b)
            a_str = str(a)
            b_str = str(b)
            a0 = int(a_str[0])
            b0 = int(b_str[0])
            a1 = int(a_str[1])
            b1 = int(b_str[1])
            if b0 == 0 or b1 == 0:
                continue
            if Decimal(a0) / Decimal(b0) == value and a1 == b1:
                nominators *= a
                denominators *= b
            if Decimal(a0) / Decimal(b1) == value and a1 == b0:
                nominators *= a
                denominators *= b
            if Decimal(a1) / Decimal(b0) == value and a0 == b1:
                nominators *= a
                denominators *= b
            if Decimal(a1) / Decimal(b1) == value and a0 == b0:
                nominators *= a
                denominators *= b

    denominators //= gcd(nominators, denominators)
    print(denominators)


euler33()
