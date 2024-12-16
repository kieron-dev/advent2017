from functools import cache


with open("input11") as f:
    nums = [int(w) for w in f.read().strip().split()]

def tick(n):
    if n == 0:
        return [1]

    digits = len(str(n))
    if digits % 2 == 0:
        half_digits = digits / 2
        pow = 10 ** half_digits
        next = int(n % pow)
        return [int(n / pow), next]

    return [n * 2024]

@cache
def count(n, iterations):
    if iterations == 1:
        return len(tick(n))

    sum = 0
    for p in tick(n):
        sum += count(p, iterations-1)

    return sum

sum_a = 0
for n in nums:
    sum_a += count(n, 25)

print(sum_a)
assert sum_a == 183248

sum_b = 0
for n in nums:
    sum_b += count(n, 75)

print(sum_b)
assert sum_b == 218811774248729
