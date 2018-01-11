from euler import timer


@timer
def euler9():
    print(next(a*b*(1000-b-a) for a in range(999, 0, -1) for b in range(a, 0, -1) if a*a+b*b == (1000-b-a)*(1000-b-a)))


euler9()
