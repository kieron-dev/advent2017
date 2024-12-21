from typing import DefaultDict
import unittest


class Grid:
    directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

    def __init__(self):
        self.cells = []
        self.start = (-1, -1)
        self.end = (-1, -1)
        self.row_count = -1
        self.col_count = -1

    def load(self, filename):
        with open(filename) as f:
            self.cells = [list(l.strip()) for l in f]
        self.row_count = len(self.cells)
        self.col_count = len(self.cells[0])
        for r, row in enumerate(self.cells):
            for c, cell in enumerate(row):
                if cell == "S":
                    self.start = (r, c)
                    self.cells[r][c] = "."
                elif cell == "E":
                    self.end = (r, c)
                    self.cells[r][c] = "."

    def print(self):
        for l in self.cells:
            print("".join(l))

    def shortest_paths(self, start):
        buckets = [set() for _ in range(self.row_count * self.col_count)]
        buckets[0].add(start)
        costs = {}
        for cost, b in enumerate(buckets):
            while len(b) > 0:
                cur = b.pop()
                if cur in costs:
                    continue
                costs[cur] = cost
                for d in self.directions:
                    r, c = cur[0] + d[0], cur[1] + d[1]
                    if r < 0 or r >= self.row_count or c < 0 or c >= self.col_count:
                        continue
                    if self.cells[r][c] == "#":
                        continue
                    buckets[cost + 1].add((r, c))
        return costs

    def cheats_table(self, dist, strict=True):
        res = DefaultDict(int)
        tos = self.shortest_paths(self.start)
        fros = self.shortest_paths(self.end)
        non_cheat = tos[self.end]

        for to, to_cost in tos.items():
            if to_cost > non_cheat: 
                continue
            for r in range(to[0] - dist, to[0] + dist + 1):
                if r < 0 or r >= self.row_count:
                    continue
                for c in range(to[1] - dist, to[1] + dist + 1):
                    if c < 0 or c >= self.col_count:
                        continue
                    mh = self.manhatten(to, (r, c))
                    if strict and mh != dist:
                        continue
                    if not strict and mh > dist:
                        continue
                    if (r, c) in fros:
                        cheat = to_cost + fros[(r, c)] + mh
                        if cheat < non_cheat:
                            res[non_cheat - cheat] += 1

        return res

    def manhatten(self, a, b):
        return abs(a[0] - b[0]) + abs(a[1] - b[1])


class TestDay20(unittest.TestCase):
    def test_example_a(self):
        grid = Grid()
        grid.load("input20a")
        costs = grid.shortest_paths(grid.start)
        self.assertEqual(84, costs[grid.end])

    def test_example_a_cheats(self):
        grid = Grid()
        grid.load("input20a")
        res = grid.cheats_table(2)
        self.assertDictEqual(
            res,
            {4: 14, 2: 14, 12: 3, 10: 2, 8: 4, 6: 2, 40: 1, 64: 1, 38: 1, 36: 1, 20: 1},
        )

    def test_part_a(self):
        grid = Grid()
        grid.load("input20")
        cheats = grid.cheats_table(2)
        sum = 0
        for k, v in cheats.items():
            if k >= 100:
                sum += v
        self.assertEqual(sum, 1378)

    def test_example_b_cheats(self):
        grid = Grid()
        grid.load("input20a")
        res = grid.cheats_table(20, strict=False)
        self.assertEqual(res[50], 32)
        self.assertEqual(res[52], 31)
        self.assertEqual(res[72], 22)

    def test_part_b(self):
        grid = Grid()
        grid.load("input20")
        cheats = grid.cheats_table(20, strict=False)
        sum = 0
        for k, v in cheats.items():
            if k >= 100:
                sum += v
        self.assertEqual(sum, 975379)


if __name__ == "__main__":
    unittest.main()
