package days_test

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("12", func() {
	It("does part A", func() {
		bytes, err := ioutil.ReadFile("input12")
		Expect(err).NotTo(HaveOccurred())

		obj := map[string]interface{}{}
		Expect(json.Unmarshal(bytes, &obj)).To(Succeed())

		Expect(sum(obj, false)).To(Equal(111754))
	})

	It("does part B", func() {
		bytes, err := ioutil.ReadFile("input12")
		Expect(err).NotTo(HaveOccurred())

		obj := map[string]interface{}{}
		Expect(json.Unmarshal(bytes, &obj)).To(Succeed())

		Expect(sum(obj, true)).To(Equal(65402))
	})
})

func sum(obj interface{}, skipRed bool) int {
	s := 0

	if i, ok := obj.(float64); ok {
		return int(i)
	}

	if arr, ok := obj.([]interface{}); ok {
		for _, a := range arr {
			s += sum(a, skipRed)
		}
		return s
	}

	if obj, ok := obj.(map[string]interface{}); ok {
		for _, val := range obj {
			if str, ok := val.(string); skipRed && ok && str == "red" {
				return 0
			}
			s += sum(val, skipRed)
		}

		return s
	}

	return 0
}
