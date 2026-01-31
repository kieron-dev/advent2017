package rule_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/rule"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {
	var (
		data      io.Reader
		validator rule.Validator
	)

	JustBeforeEach(func() {
		validator = rule.NewValidator()
		validator.Load(data)
	})

	Context("non-recursive", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`)
		})

		It("can validate", func() {
			Expect(validator.IsValid(0)).To(BeTrue())
			Expect(validator.IsValid(2)).To(BeTrue())
			Expect(validator.IsValid(1)).To(BeFalse())
			Expect(validator.IsValid(3)).To(BeFalse())
			Expect(validator.IsValid(4)).To(BeFalse())
		})

		It("can get valid count", func() {
			Expect(validator.ValidCount()).To(Equal(2))
		})
	})

	Context("recursive", func() {
		BeforeEach(func() {
			data = strings.NewReader(`
42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
`)
		})

		It("can validate correctly", func() {
			validator.SetNewRules()
			Expect(validator.IsValid(1)).To(BeTrue())
			Expect(validator.IsValid(2)).To(BeTrue())
			Expect(validator.IsValid(10)).To(BeTrue())
		})

		It("can get valid count", func() {
			validator.SetNewRules()
			Expect(validator.ValidCount()).To(Equal(12))
		})
	})
})
