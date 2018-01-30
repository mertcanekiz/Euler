from euler import timer, choose


@timer
def euler15():
    grid_size = 20
    print(choose(grid_size * 2, grid_size))


euler15()
