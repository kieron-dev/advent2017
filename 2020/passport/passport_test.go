package passport_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/passport"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Passport", func() {
	var (
		manager passport.Manager
		input   io.Reader
	)

	BeforeEach(func() {
		input = strings.NewReader(`
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`)

		manager = passport.NewManager()
		manager.Load(input)
	})

	It("creates four records", func() {
		Expect(manager.Count()).To(Equal(4))
	})

	It("reports 2 valid passports", func() {
		Expect(manager.ValidCount()).To(Equal(2))
	})
})
