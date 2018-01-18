from euler import timer
from sympy import factorint


@timer
def euler47():
    i = 646
    while True:
        i += 1
        if all(len(factorint(i+j).keys()) == 4 for j in range(4)):
            print(i)
            return


euler47()
