from collections import deque

directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]


class Grid:
    def __init__(self, size) -> None:
        self.size = size + 1
        self.start = (0, 0)
        self.end = (size, size)

    def load(self, filename, limit):
        self.coords_ls = []
        self.coords = set()
        i = 0
        with open(filename) as f:
            for l in f:
                if i == limit:
                    break
                i += 1
                l = l.strip()
                x, y = l.split(",")
                self.coords.add((int(x), int(y)))
                self.coords_ls.append(l)

    def print(self):
        for y in range(self.size):
            for x in range(self.size):
                if (x, y) in self.coords:
                    print("#", end="")
                else:
                    print(".", end="")
            print()

    def shortest_path(self):
        buckets = [deque() for _ in range((self.size) ** 2)]
        buckets[0].appendleft(self.start)
        visited = set()

        for i, bucket in enumerate(buckets):
            while len(bucket) > 0:
                cur = bucket.popleft()
                if cur in visited:
                    continue
                visited.add(cur)
                if cur == self.end:
                    return i
                for d in directions:
                    next_x, next_y = cur[0] + d[0], cur[1] + d[1]
                    if (
                        next_x < 0
                        or next_x >= self.size
                        or next_y < 0
                        or next_y >= self.size
                    ):
                        continue
                    if (next_x, next_y) in self.coords:
                        continue
                    if (next_x, next_y) in visited:
                        continue
                    buckets[i + 1].append((next_x, next_y))


grid = Grid(70)
grid.load("input18", 1024)
assert grid.shortest_path() == 364

grid = Grid(70)


def binary_search(ok, bad, fn):
    while bad > ok + 1:
        t = ok + (bad - ok) // 2
        res = fn(t)
        if res is None:
            bad = t
        else:
            ok = t
    return ok


def check(t):
    grid.load("input18", t)
    return grid.shortest_path()


part_b = ""
i = binary_search(0, 3450, check)
grid.load("input18", i + 1)
part_b = grid.coords_ls[i]
assert part_b == "52,28"
