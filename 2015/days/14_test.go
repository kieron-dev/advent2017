package days_test

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type reindeer struct {
	name      string
	speed     int
	speedTime int
	restTime  int
	score     int
}

func newReindeer(name string, speed, speedTime, restTime int) reindeer {
	return reindeer{
		name:      name,
		speed:     speed,
		speedTime: speedTime,
		restTime:  restTime,
	}
}

func (r reindeer) DistanceAt(t int) int {
	runAndRest := r.speedTime + r.restTime
	wholePeriods := t / runAndRest
	d := wholePeriods * r.speed * r.speedTime

	remainder := t - wholePeriods*runAndRest
	if remainder > r.speedTime {
		remainder = r.speedTime
	}

	return d + r.speed*remainder
}

var _ = Describe("14", func() {
	It("does part A", func() {
		input, err := os.Open("input14")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		re := regexp.MustCompile(`^(\w+) can fly (\w+) km/s for (\w+) seconds, .* (\w+) seconds\.`)
		scanner := bufio.NewScanner(input)
		maxDistance := 0
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())

			speed, err := strconv.Atoi(matches[2])
			Expect(err).NotTo(HaveOccurred())
			speedTime, err := strconv.Atoi(matches[3])
			Expect(err).NotTo(HaveOccurred())
			restTime, err := strconv.Atoi(matches[4])
			Expect(err).NotTo(HaveOccurred())
			reindeer := newReindeer(matches[1], speed, speedTime, restTime)

			dist := reindeer.DistanceAt(2503)
			if dist > maxDistance {
				maxDistance = dist
			}
		}

		Expect(maxDistance).To(Equal(2640))
	})

	It("does part B", func() {
		input, err := os.Open("input14")
		Expect(err).NotTo(HaveOccurred())
		defer input.Close()

		re := regexp.MustCompile(`^(\w+) can fly (\w+) km/s for (\w+) seconds, .* (\w+) seconds\.`)
		scanner := bufio.NewScanner(input)
		reindeers := map[string]*reindeer{}
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindStringSubmatch(line)
			Expect(matches).ToNot(BeNil())

			speed, err := strconv.Atoi(matches[2])
			Expect(err).NotTo(HaveOccurred())
			speedTime, err := strconv.Atoi(matches[3])
			Expect(err).NotTo(HaveOccurred())
			restTime, err := strconv.Atoi(matches[4])
			Expect(err).NotTo(HaveOccurred())
			reindeer := newReindeer(matches[1], speed, speedTime, restTime)
			reindeers[reindeer.name] = &reindeer
		}

		for i := 1; i < 2504; i++ {
			maxReindeer := []string{}
			maxDistance := -1

			for name, reindeer := range reindeers {
				dist := reindeer.DistanceAt(i)
				if dist == maxDistance {
					maxReindeer = append(maxReindeer, name)
				}
				if dist > maxDistance {
					maxReindeer = []string{name}
					maxDistance = dist
				}
			}

			for _, name := range maxReindeer {
				reindeers[name].score++
			}
		}

		maxScore := 0
		for _, r := range reindeers {
			if r.score > maxScore {
				maxScore = r.score
			}
		}
		Expect(maxScore).To(Equal(1102))
	})
})
