from euler import timer
from math import factorial


@timer
def euler34():
    print(sum(i for i in range(3, 7 * factorial(9))
              if sum(factorial(int(c)) for c in str(i)) == i))


euler34()
