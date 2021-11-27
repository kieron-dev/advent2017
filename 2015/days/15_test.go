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
					s := scoreCookie(ingredients, a, b, c, d, false)
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
					s := scoreCookie(ingredients, a, b, c, d, true)
					if s > maxScore {
						maxScore = s
					}
				}
			}
		}

		Expect(maxScore).To(Equal(15862900))
	})
})

func scoreCookie(ingredients []ingredient, a, b, c, d int, checkCalories bool) int {
	if checkCalories {
		calories := ingredients[0].calories*a + ingredients[1].calories*b + ingredients[2].calories*c + ingredients[3].calories*d

		if calories != 500 {
			return -1
		}
	}

	capacity := ingredients[0].capacity*a + ingredients[1].capacity*b + ingredients[2].capacity*c + ingredients[3].capacity*d
	durability := ingredients[0].durability*a + ingredients[1].durability*b + ingredients[2].durability*c + ingredients[3].durability*d
	flavor := ingredients[0].flavor*a + ingredients[1].flavor*b + ingredients[2].flavor*c + ingredients[3].flavor*d
	texture := ingredients[0].texture*a + ingredients[1].texture*b + ingredients[2].texture*c + ingredients[3].texture*d

	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}

	return capacity * durability * flavor * texture
}
