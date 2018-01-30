from euler import timer


def convergent(n, frac):
    a0 = 1
    b0 = 0
    a = frac[0]
    b = 1
    for i in range(n):
        bn = frac[1][i % len(frac[1])]
        a, a0 = bn * a + a0, a
        b, b0 = bn * b + b0, b
    return a, b


@timer
def euler65():
    # e = [2; 1, 2, 1, 1, 4, 1, 1, 6, 1, ... , 1, 2k, 1, ...]
    frac = [2, []]
    for i in range(1, 34):
        frac[1].append(1)
        frac[1].append(2 * i)
        frac[1].append(1)
    print(sum(int(c) for c in str(convergent(99, frac)[0])))


euler65()
