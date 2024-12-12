from collections import deque


with open("input12") as f:
    grid = [l.strip() for l in f]

row_count = len(grid)
col_count = len(grid[0])

seen = set()
plots = []

directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
def get_plot(r, c):
    id = grid[r][c]
    queue = deque()
    queue.append((r,c))
    plot = set()

    while len(queue) > 0:
        r, c = queue.popleft()
        if (r, c) in plot:
            continue
        plot.add((r, c))
        seen.add((r, c))
        for d in directions:
            next_r, next_c = r + d[0], c + d[1]
            if next_r < 0 or next_r >= row_count or next_c < 0 or next_c >= col_count:
                continue
            if grid[next_r][next_c] != id:
                continue
            queue.append((next_r, next_c))

    return plot

def area(plot):
    return len(plot)

def perimeter(plot):
    sum = 0
    for r, c in plot:
        if r == 0 or r == row_count - 1:
            sum += 1
        if c == 0 or c == col_count - 1:
            sum += 1
        if r > 0 and grid[r-1][c] != grid[r][c]:
            sum += 1
        if r < row_count-1 and grid[r+1][c] != grid[r][c]:
            sum += 1
        if c > 0 and grid[r][c-1] != grid[r][c]:
            sum += 1
        if c < col_count-1 and grid[r][c+1] != grid[r][c]:
            sum += 1
    return sum

for r, row in enumerate(grid):
    for c, _ in enumerate(row):
        if (r, c) in seen:
            continue
        plots.append(get_plot(r, c))

sum_a = 0
for p in plots:
    print(area(p), perimeter(p))
    sum_a += area(p) * perimeter(p)

print(sum_a)
