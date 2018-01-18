from euler import timer


def is_pentagonal(x):
    test = ((24 * x + 1)**0.5 + 1) / 6.0
    return test == int(test)


@timer
def euler44():
    i, j = 0, 0
    while True:
        i += 1
        current = i * (3 * i - 1) // 2
        for j in range(i, 0, -1):
            previous = j * (3 * j - 1) // 2
            if is_pentagonal(current - previous) and is_pentagonal(current + previous):
                print(current - previous)
                return


euler44()
