from typing import DefaultDict
import unittest


def build_forest(edges):
    sets = []
    for edge in edges:
        l, r = edge
        matching_sets = [i for i, s in enumerate(sets) if l in s or r in s]
        match len(matching_sets):
            case 0:
                sets.append({l, r})
                print("new set")
            case 1:
                idx = matching_sets[0]
                sets[idx].add(l)
                sets[idx].add(r)
                print(f"adding to set {idx}")
            case 2:
                idx1 = matching_sets[0]
                idx2 = matching_sets[1]
                sets[idx1] = sets[idx1].union(sets[idx2])
                sets = sets[:idx2] + sets[idx2+1:]
                print(f"merging sets {idx1} and {idx2}")
    return sets
                
def record_edges(edges):
    res = DefaultDict(set)
    for edge in edges:
        l, r = edge
        res[l].add(r)
        res[r].add(l)
    return res



class Test23(unittest.TestCase):
    def test_part_a(self):
        with open("input23") as f:
            edges = [l.strip().split("-") for l in f]
        res = set()
        records = record_edges(edges)
        for k, v in records.items():
            if not k.startswith("t"):
                continue
            lv = list(v)
            for i in range(len(lv)-1):
                for j in range(i+1, len(lv)):
                    one = lv[i]
                    two = lv[j]
                    if two in records[one]:
                        res.add(tuple(sorted([k, one, two])))
        ans = len(res)
        self.assertEqual(ans, 1218)

    def test_part_b(self):
        with open("input23") as f:
            edges = [tuple(l.strip().split("-")) for l in f]

        records = record_edges(edges)
        n_sets = edges
        all = set(records.keys())
        while True: 
            next_sets = set()
            for s in n_sets:
                inter = all
                for m in s:
                    inter = inter.intersection(records[m])
                if len(inter) > 0:
                    next = list(s) + [inter.pop()]
                    next_sets.add(tuple(sorted(next)))

            if len(next_sets) > 0:
                n_sets = next_sets
            else:
                break
        print(",".join(n_sets.pop()))



if __name__ == "__main__":
    unittest.main()
