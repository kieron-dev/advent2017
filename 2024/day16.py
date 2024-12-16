from __future__ import annotations
from collections import deque
import heapq
from typing import DefaultDict


class CellDetails:
    def __init__(self, distance=1e7, previous=set()) -> None:
        self.distance = distance
        self.previous = previous


class Maze:
    directions = {"e": (0, 1), "w": (0, -1), "s": (1, 0), "n": (-1, 0)}
    turns = {
        "e": {"w": 2000, "s": 1000, "n": 1000},
        "w": {"e": 2000, "s": 1000, "n": 1000},
        "n": {"s": 2000, "e": 1000, "w": 1000},
        "s": {"n": 2000, "e": 1000, "w": 1000},
    }

    def __init__(self, filename) -> None:
        with open(filename) as f:
            self.grid = [list(l.strip()) for l in f]
        self._find_start_and_end()
        self.seats = set()

    def _find_start_and_end(self):
        self.start = (-1, -1)
        self.end = (-1, -1)
        for r, row in enumerate(self.grid):
            for c, cell in enumerate(row):
                if cell == "S":
                    self.start = (r, c)
                elif cell == "E":
                    self.end = (r, c)

    def print_grid(self):
        for r, row in enumerate(self.grid):
            for c, cell in enumerate(row):
                if (r, c) in self.seats:
                    print("O", end="")
                else:
                    print(cell, end="")
            print()

    def solve(self):
        visited = set()
        positions = DefaultDict(CellDetails)
        start = (*self.start, "e")
        positions[start] = CellDetails(distance=0)

        heap = [[0, start]]

        cur_cost = 0
        end_cost = 0
        solutions = []

        while len(heap) > 0:
            cur_cost, cur_pos = heapq.heappop(heap)
            if end_cost > 0 and cur_cost > end_cost:
                break

            r, c, d = cur_pos
            if (r, c) == self.end:
                if end_cost == 0:
                    end_cost = cur_cost
                solutions.append(cur_pos)

            visited.add(cur_pos)

            new_pos = (r + self.directions[d][0], c + self.directions[d][1], d)
            new_cost = cur_cost + 1
            move = [new_cost, new_pos]
            if (
                self.grid[new_pos[0]][new_pos[1]] != "#"
                and not tuple(new_pos) in visited
                and positions[new_pos].distance >= new_cost
            ):
                heapq.heappush(heap, move)
                if positions[new_pos].distance > new_cost:
                    positions[new_pos].previous = set()
                positions[new_pos].distance = new_cost
                positions[new_pos].previous.add(cur_pos)

            for direction, cost in self.turns[d].items():
                new_cost = cur_cost + cost
                move = [new_cost, (r, c, direction)]
                new_pos = move[1]
                if not new_pos in visited and positions[new_pos].distance >= new_cost:
                    heapq.heappush(heap, move)
                    if positions[new_pos].distance > new_cost:
                        positions[new_pos].previous = set()
                    positions[new_pos].distance = new_cost
                    positions[new_pos].previous.add(cur_pos)

        self.positions = positions
        self.solutions = solutions

        return end_cost

    def get_seats(self):
        seats = set()
        queue = deque(self.solutions)
        while len(queue) > 0:
            cur = queue.popleft()
            seats.add(cur[:2])
            for p in self.positions[cur].previous:
                queue.append(p)

        self.seats = seats

        return len(seats)


if __name__ == "__main__":
    maze = Maze("input16")
    # maze.print_grid()
    cost = maze.solve()
    assert cost == 135536

    seats = maze.get_seats()
    assert seats == 583
    maze.print_grid()
