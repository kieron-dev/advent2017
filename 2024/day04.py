from collections import defaultdict


with open("input04") as f:
    lines = f.readlines()

def count_xmas(chars):
    s = "".join(chars)
    return s.count("XMAS") + s.count("SAMX")

sum_a = 0
for l in lines:
    sum_a += count_xmas(l)

for l in zip(*lines):
    sum_a += count_xmas(l)

fdiag = defaultdict(list)
bdiag = defaultdict(list)

adj = 1 - len(lines)
for x in range(len(lines[0])):
    for y in range(len(lines)):
        fdiag[x+y].append(lines[y][x])
        bdiag[x-y-adj].append(lines[y][x])

for diag in fdiag.values():
    sum_a += count_xmas(diag)

for diag in bdiag.values():
    sum_a += count_xmas(diag)

assert sum_a == 2462
print(sum_a)

sum_b = 0
for r in range(len(lines)-2):
    for c in range (len(lines[0])-2):
        if lines[r+1][c+1] != "A":
            continue
        fdiag = {lines[r][c], lines[r+2][c+2]}
        bdiag = {lines[r+2][c], lines[r][c+2]}
        if fdiag == bdiag == {"M", "S"}:
            sum_b += 1

assert sum_b == 1877
print(sum_b)
