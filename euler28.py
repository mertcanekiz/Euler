from euler import timer


def corner_sum(n):
    return (2 * n + 1)**2 + (2 * n + 1)**2 - 2 * n + (2 * n + 1)**2 - 4 * n + (2 * n + 1)**2 - 6 * n


@timer
def euler28():
    print(sum(corner_sum(n) for n in range(1, 501)) + 1)


euler28()
