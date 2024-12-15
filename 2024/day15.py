import unittest

directions = {">": (0, 1), "<": (0, -1), "^": (-1, 0), "v": (1, 0)}


def can_move(grid, position, direction):
    r, c = position
    ok = False
    while grid[r][c] != ".":
        r, c = r + directions[direction][0], c + directions[direction][1]
        if r < 0 or c < 0 or r == len(grid) or c == len(grid[0]):
            break
        if grid[r][c] == "#":
            break
    else:
        ok = True
    return ok


def move(grid, position, direction):
    r0, c0 = position
    r, c = r0 + directions[direction][0], c0 + directions[direction][1]
    if grid[r][c] != ".":
        move(grid, (r, c), direction)

    assert grid[r][c] == "."
    grid[r][c] = grid[r0][c0]
    grid[r0][c0] = "."

    return (r, c)


def can_move_expanded(grid, position, direction, ignore_box=False):
    r, c = position
    cell = grid[r][c]

    if not ignore_box and direction in ["^", "v"]:
        if cell == "[":
            right = (r, c + 1)
            return can_move_expanded(
                grid, position, direction, ignore_box=True
            ) and can_move_expanded(grid, right, direction, True)
        if cell == "]":
            left = (r, c - 1)
            return can_move_expanded(
                grid, position, direction, ignore_box=True
            ) and can_move_expanded(grid, left, direction, True)
    r1, c1 = r + directions[direction][0], c + directions[direction][1]
    if grid[r1][c1] == ".":
        return True
    if grid[r1][c1] == "#":
        return False
    return can_move_expanded(grid, (r1, c1), direction)


def move_expanded(grid, position, direction):
    r0, c0 = position
    cell = grid[r0][c0]
    positions = [position]
    if direction in ["^", "v"]:
        if cell == "[":
            right = (position[0], position[1] + 1)
            positions.append(right)
        if cell == "]":
            left = (position[0], position[1] - 1)
            positions.append(left)

    next = []
    for r1, c1 in positions:
        r, c = r1 + directions[direction][0], c1 + directions[direction][1]
        next.append((r, c))
        if grid[r][c] == "#":
            return position

    for r1, c1 in positions:
        r, c = r1 + directions[direction][0], c1 + directions[direction][1]
        if grid[r][c] != ".":
            move_expanded(grid, (r, c), direction)

        if grid[r][c] != ".":
            return position
        grid[r][c] = grid[r1][c1]
        grid[r1][c1] = "."

    return next[0]


def process_moves(grid, position, moves):
    for _, m in enumerate(moves):
        if not can_move(grid, position, m):
            continue
        position = move(grid, position, m)


def process_moves_expanded(grid, position, moves):
    for _, m in enumerate(moves):
        if not can_move_expanded(grid, position, m):
            continue
        position = move_expanded(grid, position, m)


def score(grid):
    score = 0
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell in ["O", "["]:
                score += 100 * r + c
    return score


def print_grid(grid):
    for row in grid:
        for cell in row:
            print(cell, end="")
        print()


def expand(line):
    res = []
    for l in line:
        match l:
            case "#":
                res += ["#", "#"]
            case "O":
                res += ["[", "]"]
            case ".":
                res += [".", "."]
            case "@":
                res += ["@", "."]
    return res


class Test15(unittest.TestCase):
    def test_a(self):
        grid = []
        moves = []
        start = ()
        with open("input15") as f:
            for l in f:
                l = l.strip()
                if "#" in l:
                    grid.append(list(l))
                    continue
                if l == "":
                    continue
                moves += list(l)

            found = False
            for r, row in enumerate(grid):
                if found:
                    break
                for c, e in enumerate(row):
                    if e == "@":
                        start = (r, c)
                        found = True
                        break

        process_moves(grid, start, moves)
        self.assertEqual(score(grid), 1514333)

    def test_b(self):
        grid = []
        moves = []
        start = ()
        with open("input15") as f:
            for l in f:
                l = l.strip()
                if "#" in l:
                    grid.append(expand(l))
                    continue
                if l == "":
                    continue
                moves += list(l)

            found = False
            for r, row in enumerate(grid):
                if found:
                    break
                for c, e in enumerate(row):
                    if e == "@":
                        start = (r, c)
                        found = True
                        break

        process_moves_expanded(grid, start, moves)
        self.assertEqual(score(grid), 1528453)


if __name__ == "__main__":
    unittest.main()
