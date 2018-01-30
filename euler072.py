from euler import timer, totient


@timer
def euler():
    print(sum(totient(i) for i in range(2, 1000000+1)))


euler()