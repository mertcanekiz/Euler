from euler import timer


@timer
def euler56():
    print(max(sum(int(d) for d in str(a**b)) for a in range(100) for b in range(100)))


euler56()