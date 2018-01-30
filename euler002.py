from euler import timer, fib


@timer
def euler2():
    result = 0
    i = 0
    current = 0
    while current < 4e6:
        i += 1
        current = fib(i)
        if current % 2 == 0: result += current
    print(result)


euler2()