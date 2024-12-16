import re
import unittest

def position_after(position, velocity, t, width, height):
    x = (position[0] + t * velocity[0]) % width
    y = (position[1] + t * velocity[1]) % height
    return (x, y)

def score(positions, width, height):
    tl, tr, bl, br = 0, 0, 0, 0
    for x, y in positions:
        if x < width//2:
            if y < height//2:
                tl += 1
            elif y > height//2:
                bl += 1
        elif x > width//2:
            if y < height//2:
                tr += 1
            elif y > height//2:
                br += 1
    return tl,tr,bl,br

def print_grid(positions, width, height):
    for y in range(height):
        for x in range(width):
            if (x, y) in positions:
                print("#", end="")
            else:
                print(".", end="")
        print()


class Test14(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        robots = []
        with open("input14") as f:
            for l in f:
                matches = re.match(r"p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)", l)
                assert matches is not None
                robots.append(((int(matches.group(1)), int(matches.group(2))), (int(matches.group(3)), int(matches.group(4)))))
        cls.robots = robots

    def test_a(self):
        positions = [position_after(r[0], r[1], 100, 101, 103) for r in self.robots]
        counts = score(positions, 101, 103)
        s = counts[0]*counts[1]*counts[2]*counts[3]
        self.assertEqual(s, 219512160)

    @unittest.skip("this is a manual test")
    def test_b(self):
        i = 35
        while True:
            i += 101
            positions = [position_after(r[0], r[1], i, 101, 103) for r in self.robots]
            print(i)
            print_grid(positions, 101, 103)
            input()


if __name__ == "__main__":
    unittest.main()

