with open("input06") as f:
    grid = [list(l.strip()) for l in f]

for r, row in enumerate(grid):
    if "^" in row:
        c = row.index("^")
        break
else:
    raise Exception("didn't find start")

start_row = r
start_col = c

dirs = [(-1, 0), (0, 1), (1, 0), (0, -1)]
cur_dir = 0
seen = set()

cur_row = start_row
cur_col = start_col
while True:
    seen.add((cur_row, cur_col))

    next_row = cur_row + dirs[cur_dir][0]
    next_col = cur_col + dirs[cur_dir][1]

    if next_row < 0 or next_row >= len(grid) or next_col < 0 or next_col >= len(grid[0]):
        break

    if grid[next_row][next_col] == "#":
        cur_dir = (cur_dir + 1) % 4
        continue

    cur_row = next_row
    cur_col = next_col

print(len(seen))
assert len(seen) == 5145

# part b

loops = 0

row_count = len(grid)
col_count = len(grid[0])

for r, c in seen:
    grid[r][c] = "#"

    cur_dir = 0
    seen = set()

    cur_row = start_row
    cur_col = start_col
    while True:
        if (cur_row, cur_col, cur_dir) in seen:
            loops += 1
            break

        seen.add((cur_row, cur_col, cur_dir))

        next_row = cur_row + dirs[cur_dir][0]
        next_col = cur_col + dirs[cur_dir][1]

        if next_row < 0 or next_row >= row_count or next_col < 0 or next_col >= col_count:
            break

        if grid[next_row][next_col] == "#":
            cur_dir = (cur_dir + 1) % 4
            continue

        cur_row = next_row
        cur_col = next_col

    grid[r][c] = "."

print(loops)
assert loops == 1523
