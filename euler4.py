from euler import timer


@timer
def euler4():
	print(max((i*j for i in range(999, 900, -1) for j in range(i,900,-1) if str(i*j) == str(i*j)[::-1])))


euler4()