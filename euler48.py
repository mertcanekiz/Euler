from euler import timer


@timer
def euler48():
    # This is too easy with Python
    print(str(sum(i**i for i in range(1, 1001)))[-10:])


euler48()
