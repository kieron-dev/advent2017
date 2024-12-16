with open("input10") as f:
    grid = [[int(c) for c in l.strip()] for l in f]

dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]


def reachable_nines(r, c, grid):
    cur = grid[r][c]
    if cur == 9:
        return {(r, c)}

    res = set()
    for d in dirs:
        nr, nc = r + d[0], c + d[1]
        if (
            not (0 <= nr < len(grid))
            or not (0 <= nc < len(grid[0]))
            or grid[nr][nc] != cur + 1
        ):
            continue

        for nine in reachable_nines(nr, nc, grid):
            res.add(nine)

    return res


def nine_routes(r, c, grid):
    cur = grid[r][c]
    if cur == 9:
        return 1

    res = 0
    for d in dirs:
        nr, nc = r + d[0], c + d[1]
        if (
            not (0 <= nr < len(grid))
            or not (0 <= nc < len(grid[0]))
            or grid[nr][nc] != cur + 1
        ):
            continue

        res += nine_routes(nr, nc, grid)

    return res


sum_a = 0
for r, row in enumerate(grid):
    for c, h in enumerate(row):
        if h == 0:
            sum_a += len(reachable_nines(r, c, grid))

print(sum_a)
assert sum_a == 517

sum_b = 0
for r, row in enumerate(grid):
    for c, h in enumerate(row):
        if h == 0:
            sum_b += nine_routes(r, c, grid)

print(sum_b)
assert sum_b == 1116
