from euler import timer, choose


@timer
def euler53():
    print(sum(1 for n in range(1, 101) for r in range(n) if choose(n, r) > 1000000))


euler53()
