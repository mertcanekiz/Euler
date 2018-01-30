from euler import timer


@timer
def euler81():
    grid = []
    with open('data/p081_matrix.txt') as f:
        for line in f:
            grid.append([int(n) for n in line.strip('\n').split(',')])
    for i in range(len(grid) - 2, -1, -1):
        grid[len(grid) - 1][i] += grid[len(grid) - 1][i + 1]
        grid[i][len(grid) - 1] += grid[i + 1][len(grid) - 1]
    for i in range(len(grid) - 2, -1, -1):
        for j in range(len(grid) - 2, -1, -1):
            grid[i][j] += min(grid[i + 1][j], grid[i][j + 1])
    print(grid[0][0])


euler81()
