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


def process_moves(grid, position, moves):
    for _, m in enumerate(moves):
        # print_grid(grid)
        # print(m)

        if not can_move(grid, position, m):
            continue
        position = move(grid, position, m)
    # print_grid(grid)


def score(grid):
    score = 0
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell == "O":
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
        with open("input15a") as f:
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

        print_grid(grid)
        # process_moves(grid, start, moves)
        # self.assertEqual(score(grid), 1514333)


if __name__ == "__main__":
    unittest.main()
