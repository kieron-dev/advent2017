from functools import reduce
from itertools import pairwise


def safe(report):
    if report != sorted(report) and report != sorted(report, reverse=True):
        return False

    for (a, b) in pairwise(report):
        diff = abs(a - b)
        if diff < 1 or diff > 3:
            return False

    return True


def dampenedSafe(report):
    if safe(report):
        return True

    for i in range(len(report)):
        if safe(report[:i] + report[i + 1 :]):
            return True

    return False


with open("input02") as f:
    reports = [[int(s) for s in l.split()] for l in f.readlines()]

sumA = reduce(lambda sum, report: sum + (1 if safe(report) else 0), reports, 0)
print(f"part a: {sumA}")

sumB = reduce(lambda sum, report: sum + (1 if dampenedSafe(report) else 0), reports, 0)
print(f"part b: {sumB}")
