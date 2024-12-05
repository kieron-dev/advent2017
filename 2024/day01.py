from collections import Counter
from functools import reduce

left, right = [], []
with open("input01") as f:
    for line in f:
        l, r = line.split()
        left.append(int(l))
        right.append(int(r))

sum_a = reduce(lambda x, y: x + abs(y[0] - y[1]), zip(sorted(left), sorted(right)), 0)
assert sum_a == 1603498
print(f"part a: {sum_a}")

counts = Counter(right)
sum_b = reduce(lambda x, y: x + y * counts.get(y, 0), left, 0)
assert sum_b == 25574739
print(f"part b: {sum_b}")
