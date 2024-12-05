from functools import reduce
import re


with open("input03") as f:
    instructions = f.read()

sum_a = 0
muls = re.findall(r"mul\((\d+),(\d+)\)", instructions)
sum_a += reduce(lambda x, y: x + int(y[0]) * int(y[1]), muls, 0)
assert sum_a == 175615763
print(f"part a: {sum_a}")

sum_b = 0
active = True
for i, _ in enumerate(instructions):
    if instructions[i:].startswith("do()"):
        active = True
    if instructions[i:].startswith("don't()"):
        active = False
    if not active:
        continue
    match = re.match(r"^mul\((\d+),(\d+)\)", instructions[i:])
    if match:
        sum_b += int(match.group(1)) * int(match.group(2))

assert sum_b == 74361272
print(f"part b: {sum_b}")
