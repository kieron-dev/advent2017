import re
import unittest


class TestDay13(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        with open("input13") as f:
            cls._lines = [l.strip() for l in f]

    def test_a(self):
        val = cost(self._lines, 0)
        self.assertEqual(40369, val)

    def test_b(self):
        val = cost(self._lines, 10000000000000)
        self.assertEqual(72587986598368, val)


def cost(lines, add):
    # Button A: X+40, Y+38
    # Button B: X+21, Y+84
    # Prize: X=4245, Y=5634
    # lines translate to X dir: A -> p, B -> x, Prize -> r
    # lines translate to Y dir: A -> q, B -> y, Prize -> z
    sum = 0
    p, q, r, x, y, z = 0, 0, 0, 0, 0, 0
    for l in lines:
        matches = re.match(r"Button A: X\+(\d+), Y\+(\d+)", l)
        if matches:
            p = int(matches.group(1))
            x = int(matches.group(2))
        matches = re.match(r"Button B: X\+(\d+), Y\+(\d+)", l)
        if matches:
            q = int(matches.group(1))
            y = int(matches.group(2))
        matches = re.match(r"Prize: X=(\d+), Y=(\d+)", l)
        if matches:
            r = int(matches.group(1)) + add
            z = int(matches.group(2)) + add

            n, d = r * y - z * q, p * y - x * q
            a = int(n / d)
            n1 = r - p * a
            b = int(n1 / q)
            if n1 % q != 0 or n % d != 0:
                continue
            sum += 3 * a + b
    return sum


if __name__ == "__main__":
    unittest.main()
