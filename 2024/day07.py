import math


with open("input07") as f:
    tests = [[int(n) for n in l.strip().replace(":", "").split()] for l in f]


def dfs(l):
    if len(l) == 1:
        yield l[0]
        return
    for n in dfs(l[1:]):
        yield l[0] + n
        yield l[0] * n

def dfs_b(l):
    if len(l) == 1:
        yield l[0]
        return
    for n in dfs_b(l[1:]):
        yield l[0] + n
        yield l[0] * n
        c = l[0]
        c *= 10 ** int(math.log(n, 10) + 1)
        c += n
        yield c

sum_a = 0
for t in tests:
    target = t[0]
    elements = t[1:]

    elements.reverse()
    if target in dfs(elements):
        sum_a += target

print(sum_a)
assert sum_a == 850435817339

sum_b = 0
for t in tests:
    target = t[0]
    elements = t[1:]

    elements.reverse()
    if target in dfs_b(elements):
        sum_b += target

print(sum_b)
assert sum_b == 104824810233437
