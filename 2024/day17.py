import re
import unittest


class Computer:
    def __init__(self) -> None:
        self.a = 0
        self.b = 0
        self.c = 0
        self.program = []
        self.counter = 0
        self.output = []

    def load(self, filename):
        with open(filename) as f:
            for l in f:
                l = l.strip()
                match = re.match(r"Register A: (\d+)", l)
                if match is not None:
                    self.a = int(match.group(1))
                    continue
                match = re.match(r"Register B: (\d+)", l)
                if match is not None:
                    self.b = int(match.group(1))
                    continue
                match = re.match(r"Register C: (\d+)", l)
                if match is not None:
                    self.c = int(match.group(1))
                    continue
                if l == "":
                    continue
                match = re.match(r"Program: (.*)", l)
                if match is not None:
                    self.program = [int(s) for s in match.group(1).split(",")]


    def tick(self):
        opcode = self.program[self.counter]
        operand = self.program[self.counter+1]

        jumped = False
        match opcode:
            case 0:
                power = self.combo(operand)
                self.a >>= power
            case 1:
                self.b ^= operand
            case 2:
                self.b = self.combo(operand) % 8
            case 3:
                if self.a != 0:
                    jumped = True
                    self.counter = operand
            case 4:
                self.b ^= self.c
            case 5:
                self.output.append(self.combo(operand)%8)
            case 6:
                a = self.a
                power = self.combo(operand)
                self.b = a>>power
            case 7:
                a = self.a
                power = self.combo(operand)
                self.c = a>>power
        if not jumped:
            self.counter += 2

    # def hard_coded(self):
    #     # 2,4,1,7,7,5,0,3,1,7,4,1,5,5,3,0
    #     while True:
    #         # 2,4: b is lowest 3 bits of a
    #         self.b = self.a%8
    #         # 1,7: b's 3 lowest bits are flipped
    #         self.b ^= 7
    #         # 7,5: c is set to a right shifted by b (0-7)
    #         self.c = self.a>>self.b
    #         # 0,3: a is right shifted by 3
    #         self.a >>= 3
    #         # 1,7: b's 3 lowest bits are flipped back
    #         self.b ^= 7
    #         # 4,1: b is XORed with c
    #         self.b ^= self.c
    #         # 5,5: b%8 is output
    #         self.output.append(self.b%8)
    #         # 3,0: break if a is 0
    #         if self.a == 0:
    #             break

    def get_output(self):
        return ",".join([str(n) for n in self.output])

    def run(self):
        while self.counter < len(self.program):
            self.tick()
        return self.output


    def combo(self, operand):
        if operand < 4:
            return operand
        match operand:
            case 4:
                return self.a
            case 5:
                return self.b
            case 6:
                return self.c
        raise Exception("invalid combo operand", operand)

    def reset(self):
        self.a = 0
        self.b = 0
        self.c = 0
        self.counter = 0
        self.output = []


    def do_part_b(self, pos, found_a):
        if pos < 0:
            return found_a

        l = len(self.program)

        for i in range(8):
            n = (found_a<<3)|i
            self.reset()
            self.a = n
            res = self.run()
            if res == self.program[pos:l]:
                ans = self.do_part_b(pos-1, n)
                if ans is not None:
                    return ans

class Test17(unittest.TestCase):
    def test_part_a(self):
        computer = Computer()
        computer.load("input17")
        computer.run()
        self.assertEqual(computer.get_output(), "1,0,2,0,5,7,2,1,3")

    def test_part_b(self):
        computer = Computer()
        computer.load("input17")
        ans = computer.do_part_b(len(computer.program)-1, 0)
        self.assertEqual(ans, 265652340990875)

    def test_example(self):
        computer = Computer()
        computer.load("input17a")
        computer.run()
        self.assertEqual(computer.get_output(), "4,6,3,5,6,3,5,2,1,0")

    def test_inline_1(self):
        computer = Computer()
        computer.c = 9
        computer.program = [2, 6]
        computer.run()
        self.assertEqual(computer.b, 1)

    def test_inline_2(self):
        computer = Computer()
        computer.a = 10
        computer.program = [5, 0, 5, 1, 5, 4]
        computer.run()
        self.assertEqual(computer.get_output(), "0,1,2")

    def test_inline_3(self):
        computer = Computer()
        computer.a = 2024
        computer.program = [0, 1, 5, 4, 3, 0]
        computer.run()
        self.assertEqual(computer.get_output(), "4,2,5,6,7,7,7,7,3,1,0")
        self.assertEqual(computer.a, 0)

    def test_inline_4(self):
        computer = Computer()
        computer.b = 29
        computer.program = [1, 7]
        computer.run()
        self.assertEqual(computer.b, 26)

    def test_inline_5(self):
        computer = Computer()
        computer.b = 2024
        computer.c = 43690
        computer.program = [4, 0]
        computer.run()
        self.assertEqual(computer.b, 44354)


if __name__ == "__main__":
    unittest.main()

