from collections import Counter
from functools import reduce

with open("input01") as f:
    fields = f.read().split()

left = [int(s) for s in fields[0::2]]
right = [int(s) for s in fields[1::2]]

sumA = reduce(lambda x, y: x + abs(y[0] - y[1]), zip(sorted(left), sorted(right)), 0)
print(f"part a: {sumA}")

counts = Counter(right)
sumB = reduce(lambda x, y: x + y * counts.get(y, 0), left, 0)
print(f"part b: {sumB}")
