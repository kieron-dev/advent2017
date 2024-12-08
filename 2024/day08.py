from math import gcd
from typing import DefaultDict


def print_grid(grid, antinodes):
    for r, row in enumerate(grid):
        for c, entry in enumerate(row):
            if entry != ".":
                print(entry, end="")
            elif (r, c) in antinodes:
                print("#", end="")
            else:
                print(".", end="", )
        print("")

with open("input08") as f:
    grid = [line.strip() for line in f]

row_count = len(grid)
col_count = len(grid[0])

locations = DefaultDict(set)
for r, line in enumerate(grid):
    for c, entry in enumerate(line):
        if entry == ".":
            continue
        locations[entry].add((r, c))

antinodes = set()

for freq_locations in [list(vals) for vals in locations.values()]:
    for i in range(len(freq_locations)):
        for j in range(i+1, len(freq_locations)):
            antenna_1 = freq_locations[i]
            antenna_2 = freq_locations[j]
            diff_r = antenna_1[0] - antenna_2[0]
            diff_c = antenna_1[1] - antenna_2[1]
            antinode_1 = (antenna_1[0]+diff_r, antenna_1[1]+diff_c)
            antinode_2 = (antenna_2[0]-diff_r, antenna_2[1]-diff_c)
            if 0 <= antinode_1[0] < row_count and 0 <= antinode_1[1] < col_count:
                antinodes.add(antinode_1)
            if 0 <= antinode_2[0] < row_count and 0 <= antinode_2[1] < col_count:
                antinodes.add(antinode_2)

# print_grid(grid, antinodes)

print(len(antinodes))
assert len(antinodes) == 400

antinodes = set()

for freq_locations in [list(vals) for vals in locations.values()]:
    for i in range(len(freq_locations)):
        for j in range(i+1, len(freq_locations)):
            antenna_1 = freq_locations[i]
            antenna_2 = freq_locations[j]
            diff_r = antenna_1[0] - antenna_2[0]
            diff_c = antenna_1[1] - antenna_2[1]
            gcd_r_c = gcd(diff_r, diff_c)
            diff_r /= gcd_r_c
            diff_c /= gcd_r_c
            antinode = antenna_1
            while 0 <= antinode[0] < row_count and 0 <= antinode[1] < col_count:
                antinodes.add(antinode)
                antinode = (antinode[0]-diff_r, antinode[1]-diff_c)
            antinode = antenna_1
            while 0 <= antinode[0] < row_count and 0 <= antinode[1] < col_count:
                antinodes.add(antinode)
                antinode = (antinode[0]+diff_r, antinode[1]+diff_c)

# print_grid(grid, antinodes)

print(len(antinodes))
assert(len(antinodes)) == 1280
