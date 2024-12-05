from collections import defaultdict
from functools import cmp_to_key


rules = defaultdict(set)
updates = []
with open("input05") as f:
    for line in f:
        line = line.strip()
        if line == "":
            continue

        if "|" in line:
            before, after = line.split("|")
            rules[before].add(after)
            continue

        updates.append(line.split(","))

sum_a: int = 0
bad_updates: list[str] = []
for update in updates:
    seen = set()
    res = []
    for page in update:
        for after in rules[page]:
            if after in seen:
                break
        else:
            res.append(page)

        seen.add(page)

    if res == update:
        sum_a += int(update[len(update)//2])
    else:
        bad_updates.append(update)

assert sum_a == 5639
print(sum_a)

def cmp(l: str, r: str):
    if l in rules[r]:
        return 1
    if r in rules[l]:
        return -1
    return 0

sum_b = 0
for update in bad_updates:
    fixed = sorted(update, key=cmp_to_key(cmp))
    sum_b += int(fixed[len(fixed)//2])

assert sum_b == 5273
print(sum_b)

