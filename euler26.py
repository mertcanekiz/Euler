from euler import timer


def decimal_period(n):
    while n%2 == 0:
        n //= 2
    while n%5 == 0:
        n //= 5
    if n == 1:
        return 0
    result = 1
    while True:
        current = 10**result % n
        if current == 1:
            return result
        result += 1
    return result


@timer
def euler26():
    max_period, result = 0, 0
    for i in range(3, 10):
        current = decimal_period(i)
        print(i, current)
        if current > max_period:
            max_period, result = current, i
    print(result, max_period)


euler26()
