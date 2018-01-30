from euler import timer
import networkx


@timer
def euler83():
    grid = []
    with open('data/p081_matrix.txt') as f:
        for line in f:
            grid.append([int(n) for n in line.strip('\n').split(',')])

    n, m = len(grid), len(grid[0])
    graph = networkx.DiGraph()
    for i in range(n):
        for j in range(m):
            neighbors = [(i + x, j + y) for x, y in ((-1, 0), (0, -1), (1, 0), (0, 1))
                         if 0 <= i + x < n and 0 <= j + y < m]
            for nx, ny in neighbors:
                graph.add_edge((i, j), (nx, ny), weight=grid[nx][ny])

    path_length = networkx.dijkstra_path_length(graph, source=(0, 0), target=(n - 1, m - 1))
    print(grid[0][0] + path_length)


euler83()
