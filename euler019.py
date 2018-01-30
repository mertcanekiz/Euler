from euler import timer
from datetime import date, timedelta


@timer
def euler19():
    start = date(1901, 1, 1)
    end = date(2000, 12, 31)
    current = start
    result = 0
    while current < end:
        if current.day == 1 and current.weekday() == 6:
            result += 1
        current += timedelta(days=1)
    print(result)


euler19()
