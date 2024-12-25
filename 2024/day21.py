import unittest


class Keypad:
    layout = {}

    def __init__(self):
        self.pos = self.layout["A"]
        self.state = {}
        for key, coord in self.layout.items():
            self.state[coord] = key

    def move(self, key):
        self.pos = self.layout[key]

    def get_directions(self, _):
        return []


class NumericKeypad(Keypad):
    layout = {
        "A": (3, 2),
        "0": (3, 1),
        "1": (2, 0),
        "2": (2, 1),
        "3": (2, 2),
        "4": (1, 0),
        "5": (1, 1),
        "6": (1, 2),
        "7": (0, 0),
        "8": (0, 1),
        "9": (0, 2),
    }

    def get_directions(self, target):
        to = self.layout[target]

        p = list(self.pos)
        option_2 = []
        while p[1] > to[1]:
            p[1] -= 1
            option_2.append("<")
        while p[1] < to[1]:
            p[1] += 1
            option_2.append(">")
        while p[0] > to[0]:
            p[0] -= 1
            option_2.append("^")
        while p[0] < to[0]:
            p[0] += 1
            option_2.append("v")

        p = list(self.pos)
        option_1 = []
        while p[0] > to[0]:
            p[0] -= 1
            option_1.append("^")
        while p[0] < to[0]:
            p[0] += 1
            option_1.append("v")
        while p[1] > to[1]:
            p[1] -= 1
            option_1.append("<")
        while p[1] < to[1]:
            p[1] += 1
            option_1.append(">")

        res = []
        if self.pos[1] != 0 or to[0] != 3:
            res.append(option_1 + ["A"])
        if self.pos[0] != 3 or to[1] != 0:
            res.append(option_2 + ["A"])

        return res


class CursorKeypad(Keypad):
    layout = {
        "A": (0, 2),
        "^": (0, 1),
        "<": (1, 0),
        "v": (1, 1),
        ">": (1, 2),
    }

    def get_directions(self, target):
        to = self.layout[target]
        res = []

        p = list(self.pos)
        if self.pos[0] != 0 or to[1] != 0:
            option_1 = []
            while p[1] > to[1]:
                p[1] -= 1
                option_1.append("<")
            while p[1] < to[1]:
                p[1] += 1
                option_1.append(">")
            while p[0] < to[0]:
                p[0] += 1
                option_1.append("v")
            while p[0] > to[0]:
                p[0] -= 1
                option_1.append("^")
            option_1.append("A")
            # return [option_1]
            res.append(option_1)

        p = list(self.pos)
        if self.pos[1] != 0 or to[0] != 0:
            option_2 = []
            while p[0] < to[0]:
                p[0] += 1
                option_2.append("v")
            while p[0] > to[0]:
                p[0] -= 1
                option_2.append("^")
            while p[1] > to[1]:
                p[1] -= 1
                option_2.append("<")
            while p[1] < to[1]:
                p[1] += 1
                option_2.append(">")
            option_2.append("A")
            # return [option_2]
            res.append(option_2)

        return res


class Chain:
    def __init__(self, cursor_count):
        self.cursor_count = cursor_count
        self.cache = {}

    def load(self, filename):
        with open(filename) as f:
            self.codes = [l.strip() for l in f]

    def propagate(self, keys, level=0):
        cache_key = (tuple(keys), level)
        if cache_key in self.cache:
            return self.cache[cache_key]

        if level == self.cursor_count + 1:
            return len(keys)

        res = 0
        if level == 0:
            keypad = NumericKeypad()
        else:
            keypad = CursorKeypad()
        for k in keys:
            best = 0
            options = keypad.get_directions(k)
            for option in options:
                next = self.propagate(option, level + 1)
                if best == 0 or next < best:
                    best = next
            keypad.move(k)
            res += best

        self.cache[cache_key] = res
        return res

    def solve_one(self, input):
        res = self.propagate(input)
        num = int(input[:-1])
        return res * num

    def solve(self):
        sum = 0
        for code in self.codes:
            sum += self.solve_one(code)
        return sum


class TestDay21(unittest.TestCase):
    def test_example_a(self):
        self.assertEqual(Chain(2).solve_one("029A"), 68 * 29)
        self.assertEqual(Chain(2).solve_one("980A"), 60 * 980)
        self.assertEqual(Chain(2).solve_one("179A"), 68 * 179)
        self.assertEqual(Chain(2).solve_one("456A"), 64 * 456)
        self.assertEqual(Chain(2).solve_one("379A"), 64 * 379)

    def test_part_a(self):
        chain = Chain(2)
        chain.load("input21")
        self.assertEqual(chain.solve(), 163920)

    def test_part_b(self):
        chain = Chain(25)
        chain.load("input21")
        self.assertEqual(chain.solve(), 204040805018350)


if __name__ == "__main__":
    unittest.main()
