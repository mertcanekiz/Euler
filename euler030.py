from euler import timer


@timer
def euler30():
    limit = 10*9**5
    print(sum(i for i in range(2, limit) if sum(int(c)**5 for c in str(i)) == i))


euler30()
