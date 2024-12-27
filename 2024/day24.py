import unittest


class Gate:
    def __init__(self, l, r, op):
        self.op = op
        self.l = l
        self.r = r

    def eval(self, inputs):
        print(f"{self.l} {self.op} {self.r}")
        match self.op:
            case "XOR":
                return inputs[self.l] ^ inputs[self.r]
            case "OR":
                return inputs[self.l] | inputs[self.r]
            case "AND":
                return inputs[self.l] & inputs[self.r]
        raise Exception("oops")

class Circuit:
    def __init__(self):
        self.inputs = {}
        self.gates = {}
        pass

    def load(self, filename):
        with open(filename) as f:
            for l in f:
                l = l.strip()
                if ":" in l:
                    l, r = l.split(": ")
                    self.inputs[l] = int(r)
                    continue

                if l == "":
                    continue

                l, op, r, _, out = l.split()
                self.gates[out] = Gate(l, r, op)
    def swap(self, a, b):
        old = self.gates[a]
        self.gates[a] = self.gates[b]
        self.gates[b] = old

    def evaluate(self, out):
        if out in self.inputs:
            return self.inputs[out]

        gate = self.gates[out]
        for input in [gate.l, gate.r]:
            self.evaluate(input)

        print(out, end=" = ")
        self.inputs[out] = gate.eval(self.inputs)
        return self.inputs[out]

class Test24(unittest.TestCase):
    def test_part_a(self):
        circuit = Circuit()
        circuit.load("input24")
        res = 0
        for i in range(46):
            res += (circuit.evaluate(f"z{i:02d}") << i)
        self.assertEqual(res, 56278503604006)

    def test_bits(self):
        circuit = Circuit()
        circuit.load("input24")
        circuit.swap("dhg", "z06")
        circuit.swap("nbf", "z38")
        circuit.swap("bhd", "z23")
        circuit.swap("brk", "dpd")
        for i in range(46):
            out = f"z{i:02d}"
            circuit.evaluate(out)
            print()

    def test_part_b(self):
        swaps = ["dhg", "z06", "nbf", "z38", "bhd", "z23", "brk", "dpd"]
        print(",".join(sorted(swaps)))





if __name__ == "__main__":
    unittest.main()
