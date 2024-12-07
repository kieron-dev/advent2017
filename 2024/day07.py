import math


with open("input07") as f:
    tests = [[int(n) for n in l.strip().replace(":", "").split()] for l in f]

sum_a = 0
for t in tests:
    target = t[0]
    elements = t[1:]

    for p in range(1 << (len(elements)-1)):
        s = elements[0]
        for i, e in enumerate(elements[1:]):
            if p & 1<<i:
                s += e
            else:
                s *= e
            if s > target:
                break
        if s == target:
            sum_a += target
            break

print(sum_a)
assert sum_a == 850435817339

sum_b = 0
for t in tests:
    target = t[0]
    elements = t[1:]

    for p in range(3**(len(elements)-1)):
        s = elements[0]
        b3 = p
        for i, e in enumerate(elements[1:]):
            match b3 % 3:
                case 0:
                    s += e
                case 1:
                    s *= e
                case 2:
                    s *= 10**int(math.log(e, 10)+1)
                    s += e
            if s > target:
                break
            b3 //= 3
        if s == target:
            sum_b += target
            break

print(sum_b)
assert sum_b == 104824810233437
