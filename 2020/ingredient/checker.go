// Package ingredient checks for allergens
package ingredient

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strings"
)

type Allergen struct {
	name       string
	ingredient *Ingredient
}

func NewAllergen(name string) *Allergen {
	return &Allergen{
		name: name,
	}
}

type Ingredient struct {
	name     string
	allergen *Allergen
}

func NewIngredient(name string) *Ingredient {
	return &Ingredient{
		name: name,
	}
}

func (i *Ingredient) SetAllergen(a *Allergen) {
	i.allergen = a
	a.ingredient = i
}

type Product struct {
	ingredients []*Ingredient
	allergens   []*Allergen
}

type Checker struct {
	products    []*Product
	ingredients map[string]*Ingredient
	allergens   map[string]*Allergen
	counts      map[string]int
}

func NewChecker() Checker {
	return Checker{
		ingredients: map[string]*Ingredient{},
		allergens:   map[string]*Allergen{},
		counts:      map[string]int{},
	}
}

func (c *Checker) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		prod := Product{}

		split1 := strings.Split(line, " (contains ")
		for _, ingrStr := range strings.Split(split1[0], " ") {
			if c.ingredients[ingrStr] == nil {
				c.ingredients[ingrStr] = NewIngredient(ingrStr)
			}
			prod.ingredients = append(prod.ingredients, c.ingredients[ingrStr])
			c.counts[ingrStr]++
		}

		for _, allergenStr := range strings.Split(split1[1][:len(split1[1])-1], ", ") {
			if c.allergens[allergenStr] == nil {
				c.allergens[allergenStr] = NewAllergen(allergenStr)
			}
			prod.allergens = append(prod.allergens, c.allergens[allergenStr])
		}

		c.products = append(c.products, &prod)
	}
}

func (c *Checker) FindImpossible() []*Ingredient {
	impossible := []*Ingredient{}

	for _, i := range c.ingredients {
		possibleAllergens := map[*Allergen]bool{}

		for _, prod := range c.products {
			if containsIngredient(prod.ingredients, i) {
				for _, allergen := range prod.allergens {
					possibleAllergens[allergen] = true
				}
			}
		}

		for _, prod := range c.products {
			if !containsIngredient(prod.ingredients, i) {
				for _, allergen := range prod.allergens {
					delete(possibleAllergens, allergen)
				}
			}
		}

		if len(possibleAllergens) == 0 {
			impossible = append(impossible, i)
		}
	}

	return impossible
}

func (c *Checker) assignAllergens() {
	toFind := len(c.allergens)

	for toFind > 0 {
		somethingChanged := false

		for _, a := range c.allergens {
			if a.ingredient != nil {
				continue
			}

			ingredients := []*Ingredient{}
			for _, i := range c.ingredients {
				ingredients = append(ingredients, i)
			}

			for _, p := range c.products {
				if containsAllergen(p.allergens, a) {
					ingredients = ingredientIntersection(ingredients, p.ingredients)
				}
			}

			if len(ingredients) == 1 {
				ingredients[0].SetAllergen(a)
				c.RemoveAllergen(a)
				c.RemoveIngredient(ingredients[0])
				toFind--
				somethingChanged = true
				break
			}
		}

		if !somethingChanged {
			log.Fatalf("need a better strategy")
		}
	}
}

func (c *Checker) RemoveIngredient(i *Ingredient) {
	for _, prod := range c.products {
		prod.RemoveIngredient(i)
	}
}

func (p *Product) RemoveIngredient(i *Ingredient) {
	newIngredients := []*Ingredient{}

	for _, o := range p.ingredients {
		if o == i {
			continue
		}

		newIngredients = append(newIngredients, o)
	}

	p.ingredients = newIngredients
}

func (c *Checker) RemoveAllergen(a *Allergen) {
	for _, prod := range c.products {
		prod.RemoveAllergen(a)
	}
}

func (p *Product) RemoveAllergen(a *Allergen) {
	newAllergens := []*Allergen{}

	for _, o := range p.allergens {
		if o == a {
			continue
		}

		newAllergens = append(newAllergens, o)
	}

	p.allergens = newAllergens
}

func containsIngredient(ingredients []*Ingredient, needle *Ingredient) bool {
	for _, a := range ingredients {
		if a == needle {
			return true
		}
	}

	return false
}

func containsAllergen(allergens []*Allergen, needle *Allergen) bool {
	for _, a := range allergens {
		if a == needle {
			return true
		}
	}

	return false
}

func ingredientIntersection(ingrA, ingrB []*Ingredient) []*Ingredient {
	m := map[*Ingredient]bool{}
	res := []*Ingredient{}

	for _, i := range ingrA {
		m[i] = true
	}

	for _, i := range ingrB {
		if m[i] {
			res = append(res, i)
		}
	}

	return res
}

func (c *Checker) NonAllergenCount() int {
	num := 0
	for _, i := range c.FindImpossible() {
		num += c.counts[i.name]
	}

	return num
}

func (c *Checker) AllergenIngredients() string {
	for _, i := range c.FindImpossible() {
		c.RemoveIngredient(i)
		delete(c.ingredients, i.name)
	}

	c.assignAllergens()

	ingredients := []*Ingredient{}
	for _, i := range c.ingredients {
		ingredients = append(ingredients, i)
	}

	sort.Slice(ingredients, func(a, b int) bool {
		return ingredients[a].allergen.name < ingredients[b].allergen.name
	})

	s := ""
	for _, i := range ingredients {
		s += i.name + ","
	}

	return s[:len(s)-1]
}
