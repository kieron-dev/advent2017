with open("input04") as f:
    lines = [[c for c in line.strip()] for line in f ]

def countXMAS(chars):
    c = 0
    s = "".join(chars)
    c += s.count("XMAS")
    s = "".join(list(reversed(chars)))
    c += s.count("XMAS")
    return c

sumA = 0
for l in lines:
    sumA += countXMAS(l)

for l in zip(*lines):
    sumA += countXMAS(l)

max_col = len(lines[0])
max_row = len(lines)
fdiag = [[] for _ in range(max_row + max_col - 1)]
bdiag = [[] for _ in range(len(fdiag))]
adj = -max_row + 1

for x in range(max_col):
    for y in range(max_row):
        fdiag[x+y].append(lines[y][x])
        bdiag[x-y-adj].append(lines[y][x])


for l in fdiag:
    sumA += countXMAS(l)
for l in bdiag:
    sumA += countXMAS(l)

assert sumA == 2462
print(sumA)

sumB = 0
for r in range(len(lines)-2):
    for c in range (len(lines[0])-2):
        if lines[r+1][c+1] != "A":
            continue
        d1 = {lines[r][c], lines[r+2][c+2]}
        d2 = {lines[r+2][c], lines[r][c+2]}
        if d1 == {"M", "S"} and d2 == {"M", "S"}:
            sumB += 1

assert sumB == 1877
print(sumB)
