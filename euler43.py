from euler import timer
from itertools import permutations


@timer
def euler43():
    total = 0
    for perm in permutations('0123456789'):
        current = ''.join(perm)
        if int(current[7:10]) % 17 == 0\
            and int(current[6:9]) % 13 == 0\
            and int(current[5:8]) % 11 == 0\
            and int(current[4:7]) % 7 == 0\
            and int(current[3:6]) % 5 == 0\
            and int(current[2:5]) % 3 == 0\
            and int(current[1:4]) % 2 == 0:
            total += int(current)
    print(total)


euler43()
