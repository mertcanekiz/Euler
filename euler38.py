from euler import timer, is_pandigital


def concatenated_product(x, n):
    return ''.join(str(x*i) for i in range(1, n + 1))


@timer
def euler38():
    print(max(int(concatenated_product(x, 2)) for x in range(9123, 9876) if is_pandigital(concatenated_product(x, 2))))


euler38()
