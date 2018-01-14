from euler import timer, is_pandigital


@timer
def euler32():
    print(sum(set(a * b for a in range(2, 60) for b in range(1234 if a < 10 else 123, 10000 // a) if is_pandigital(str(a) + str(b) + str(a * b)))))


euler32()
