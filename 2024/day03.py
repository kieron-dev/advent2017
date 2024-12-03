from functools import reduce
import re


with open("input03") as f:
    instructions = f.read()

sumA = 0
muls = re.findall(r"mul\((\d+),(\d+)\)", instructions)
sumA += reduce(lambda x, y: x + int(y[0]) * int(y[1]), muls, 0)
assert sumA == 175615763
print(f"part a: {sumA}")

sumB = 0
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
        sumB += int(match.group(1)) * int(match.group(2))

print(f"part b: {sumB}")
assert sumB == 74361272
