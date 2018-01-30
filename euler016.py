from euler import timer


@timer
def euler16():
    print(sum(int(c) for c in str(2**1000)))


euler16()
