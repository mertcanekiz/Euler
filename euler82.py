from euler import timer


@timer
def euler82():
    grid = []
    with open('data/p082_matrix.txt') as f:
        for line in f:
            grid.append([int(n) for n in line.strip('\n').split(',')])
    n, m = len(grid), len(grid[0])
    cost = [grid[i][-1] for i in range(n)]
    for i in range(m - 2, -1, -1):
        cost[0] += grid[0][i]
        for j in range(1, n):
            cost[j] = min(cost[j], cost[j-1]) + grid[j][i]
        for j in range(n - 2, -1, -1):
            cost[j] = min(cost[j], cost[j+1] + grid[j][i])
    
    print(min(cost))


euler82()
