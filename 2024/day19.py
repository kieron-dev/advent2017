import unittest


class Towels:
    def __init__(self, filename) -> None:
        self.available = set()
        self.wanted = []
        self.cache = {}

        with open(filename) as f:
            for l in f:
                l = l.strip()
                if l == "":
                    continue
                if "," in l:
                    self.available = set(l.split(", "))
                    continue

                self.wanted.append(l)

    def is_possible(self, pattern):
        if pattern == "":
            return True
        for a in self.available:
            if pattern.startswith(a) and self.is_possible(pattern[len(a) :]):
                return True
        return False

    def count_ways(self, pattern):
        if pattern in self.cache:
            return self.cache[pattern]

        if pattern == "":
            self.cache[pattern] = 1
            return 1
        c = 0
        for a in self.available:
            if pattern.startswith(a):
                c += self.count_ways(pattern[len(a) :])
        self.cache[pattern] = c
        return c

    def count_possible(self):
        c = 0
        for p in self.wanted:
            if self.is_possible(p):
                c += 1
        return c

    def sum_all_ways(self):
        c = 0
        for p in self.wanted:
            c += self.count_ways(p)
        return c


class Test19(unittest.TestCase):
    def test_example_a(self):
        towels = Towels("input19a")
        self.assertEqual(towels.count_possible(), 6)

    def test_example_b(self):
        towels = Towels("input19a")
        self.assertEqual(towels.sum_all_ways(), 16)

    def test_part_a(self):
        towels = Towels("input19")
        self.assertEqual(towels.count_possible(), 333)

    def test_part_b(self):
        towels = Towels("input19")
        self.assertEqual(towels.sum_all_ways(), 678536865274732)


if __name__ == "__main__":
    unittest.main()
