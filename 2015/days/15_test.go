package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var re = regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func newIngredient(line string) ingredient {
	matches := re.FindStringSubmatch(line)
	Expect(matches).ToNot(BeNil())

	return ingredient{
		name:       matches[1],
		capacity:   toInt(matches[2]),
		durability: toInt(matches[3]),
		flavor:     toInt(matches[4]),
		texture:    toInt(matches[5]),
		calories:   toInt(matches[6]),
	}
}

func toInt(a string) int {
	n, err := strconv.Atoi(a)
	Expect(err).NotTo(HaveOccurred())
	return n
}

var _ = Describe("15", func() {
	It("does part A", func() {
		input, err := os.Open("input15")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		ingredients := []ingredient{}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			ingredient := newIngredient(scanner.Text())
			ingredients = append(ingredients, ingredient)
		}

		Expect(ingredients).To(HaveLen(4))

		maxScore := 0
		for a := 0; a <= 100; a++ {
			for b := 0; a+b <= 100; b++ {
				for c := 0; a+b+c <= 100; c++ {
					d := 100 - a - b - c
					s := scoreCookie(ingredients, []int{a, b, c, d}, false)
					if s > maxScore {
						maxScore = s
					}
				}
			}
		}

		Expect(maxScore).To(Equal(18965440))
	})

	It("does part B", func() {
		input, err := os.Open("input15")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		ingredients := []ingredient{}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			ingredient := newIngredient(scanner.Text())
			ingredients = append(ingredients, ingredient)
		}

		Expect(ingredients).To(HaveLen(4))

		maxScore := 0
		for a := 0; a <= 100; a++ {
			for b := 0; a+b <= 100; b++ {
				for c := 0; a+b+c <= 100; c++ {
					d := 100 - a - b - c
					s := scoreCookie(ingredients, []int{a, b, c, d}, true)
					if s > maxScore {
						maxScore = s
					}
				}
			}
		}

		Expect(maxScore).To(Equal(15862900))
	})
})

func scoreCookie(ingredients []ingredient, amounts []int, checkCalories bool) int {
	if checkCalories {
		calories := sumProperty(ingredients, amounts, func(i ingredient) int {
			return i.calories
		})

		if calories != 500 {
			return -1
		}
	}

	capacity := sumProperty(ingredients, amounts, func(i ingredient) int {
		return i.capacity
	})

	durability := sumProperty(ingredients, amounts, func(i ingredient) int {
		return i.durability
	})

	flavor := sumProperty(ingredients, amounts, func(i ingredient) int {
		return i.flavor
	})

	texture := sumProperty(ingredients, amounts, func(i ingredient) int {
		return i.texture
	})

	return capacity * durability * flavor * texture
}

func sumProperty(
	ingredients []ingredient,
	amounts []int,
	getter func(ingred ingredient) int,
) int {

	property := 0
	for i, ingredient := range ingredients {
		property += getter(ingredient) * amounts[i]
	}
	if property < 0 {
		property = 0
	}

	return property
}
