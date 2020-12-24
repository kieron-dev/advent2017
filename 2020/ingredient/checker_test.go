package ingredient_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/ingredient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checker", func() {
	var (
		data    io.Reader
		checker ingredient.Checker
	)

	BeforeEach(func() {
		data = strings.NewReader(`
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`)

		checker = ingredient.NewChecker()
		checker.Load(data)
	})

	It("finds the non-allergens", func() {
		Expect(checker.NonAllergenCount()).To(Equal(5))
	})

	It("can order the allergen ingredients", func() {
		Expect(checker.AllergenIngredients()).To(Equal("mxmxvkd,sqjhc,fvjkl"))
	})
})
