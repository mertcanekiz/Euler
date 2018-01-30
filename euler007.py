from euler import timer, isprime


@timer
def euler7():
    numprimes = 1
    counter = 1
    limit = 10001
    while numprimes < limit:
        counter += 2
        if isprime(counter):
            numprimes += 1
    print(counter)


euler7()
