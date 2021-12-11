package days_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func calcPrimes(upto int) []int {
	sieve := make([]bool, upto)
	sieve[1] = true
	sieve[2] = true

	primes := []int{2}

	for i := 3; i < upto; i++ {
		if !sieve[i] {
			primes = append(primes, i)
			for j := i; j < upto; j += i {
				sieve[j] = true
			}
		}
	}

	return primes
}

var _ = Describe("20", func() {
	It("calcs the sum of divisors", func() {
		primes := calcPrimes(1000)
		Expect(sumOfDivisors(1, primes)).To(Equal(1))
		Expect(sumOfDivisors(2, primes)).To(Equal(3))
		Expect(sumOfDivisors(3, primes)).To(Equal(4))
		Expect(sumOfDivisors(4, primes)).To(Equal(7))
	})

	FIt("does part A", func() {
		lim := 36_000_000
		primeLim := 1_000_000

		primes := calcPrimes(primeLim)

		i := 1
		for {
			presents := 10 * sumOfDivisors(i, primes)
			if i%10000 == 0 {
				fmt.Printf("%12d: %12d\n", i, presents)
			}
			if presents >= lim {
				break
			}
			i++
			if i >= primeLim {
				Fail("need a higher prime limit")
			}
		}

		Expect(i).To(Equal(831600))
	})
})

func pow(n, p int) int {
	r := 1
	for i := 0; i < p; i++ {
		r *= n
	}

	return r
}

func sumOfDivisors(n int, primes []int) int {
	sum := 1

	pn := 0
	for n > 1 {
		p := primes[pn]
		i := 0
		for n%p == 0 {
			i++
			n /= p
		}
		if i > 0 {
			sum *= (pow(p, i+1) - 1) / (p - 1)
		}
		pn++
	}

	return sum
}
