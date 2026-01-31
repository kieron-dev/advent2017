package ingredient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIngredient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ingredient Suite")
}
