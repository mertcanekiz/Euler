from euler import timer
from math import sqrt


@timer
def euler85():
    target = 2000000
    area = 0
    diff = float('Inf')
    x, y = 2000, 1
    while x >= y:
        current = x * (x + 1) * y * (y + 1) // 4
        if diff > abs(current - target):
            diff, area = abs(current - target), x * y
        y += 1
        x = int((4 * target / (y * y + y))**0.5)
    print(area)


euler85()
