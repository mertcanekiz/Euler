from euler import timer, lcm
from functools import reduce


@timer
def euler5():
	print(reduce(lcm, range(1, 21)))


euler5()