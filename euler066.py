from euler import timer, continued_fraction, convergent


def pell(d):
    i = 0
    frac = continued_fraction(d)
    while True:
        i += 1
        c = convergent(i, frac)
        if c[0]**2 - d*c[1]**2 == 1:
            return c


@timer
def euler66():
    max_x, max_d = 0, 0
    for d in range(2, 1000):
        if int(d**0.5) == d**0.5: continue
        current = pell(d)
        if current[0] > max_x:
            max_x, max_d = current[0], d
    print(max_x, max_d)


euler66()
