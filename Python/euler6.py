from euler import timer


@timer
def euler6():
    print(sum(range(1, 101))**2-sum(i*i for i in range(1, 101)))


euler6()
