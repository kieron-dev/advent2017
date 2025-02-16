from functools import cmp_to_key
import unittest


class Secret:
    def __init__(self, num):
        self.secret = num
        self.prune = (1<<24) - 1

    def next(self):
        n = self.secret
        p = n
        n <<= 6
        n ^= p
        n &= self.prune

        p = n
        n >>= 5
        n ^= p
        n &= self.prune

        p =n 
        n <<= 11
        n ^= p
        n &= self.prune

        self.secret = n
        return n

    def next_2000(self):
        n = 0
        for _ in range(2000):
            n = self.next()
        return n

    def seq_2000(self):
        p = self.secret%10
        s = []
        v = []
        for _ in range(2000):
            n = self.next() % 10
            d = n - p
            s.append(d)
            v.append(n)
            p = n
        return s, v

    def freqs(self):
        seq, vals = self.seq_2000()
        seqs = {}
        for i, _ in enumerate(seq[:-3]):
            key = tuple(seq[i:i+4])
            if not key in seqs:
                seqs[key] = vals[i+3]
        return seqs
            
        
class Test22(unittest.TestCase):
    def test_example_a(self):
        secret = Secret(123)
        self.assertEqual(secret.next(), 15887950)
        self.assertEqual(secret.next(), 16495136)
        self.assertEqual(secret.next(), 527345)

    def test_2000_times(self):
        secret = Secret(2024)
        n = 0
        for _ in range(2000):
            n = secret.next()
        self.assertEqual(n, 8667524)

    def test_part_a(self):
        sum = 0
        with open("input22") as f:
            for l in f:
                secret = Secret(int(l.strip()))
                sum += secret.next_2000()
        self.assertEqual(sum , 14623556510)

    def test_part_b(self):
        freqs = []
        with open("input22") as f:
            for l in f:
                secret = Secret(int(l.strip()))
                freqs.append(secret.freqs())

        sums = {}
        for freq in freqs:
            for k, v in freq.items():
                if k not in sums:
                    sums[k] = 0
                sums[k] += v

        res = dict(sorted(sums.items(), key=lambda item: item[1], reverse=True))
        v = next(iter(res.values()))
        print(v)
        self.assertEqual(1701, v)



if __name__ == "__main__":
    unittest.main()

        
