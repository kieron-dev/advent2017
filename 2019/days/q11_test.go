package days_test

import (
	"io/ioutil"
	"strings"

	"github.com/kieron-pivotal/advent2017/advent2019/grid"
	"github.com/kieron-pivotal/advent2017/advent2019/robot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Q11", func() {
	var (
		r    *robot.Robot
		prog string
	)

	BeforeEach(func() {
		r = robot.New()
		progBytes, err := ioutil.ReadFile("../days/input11")
		if err != nil {
			panic(err)
		}
		prog = strings.TrimSpace(string(progBytes))
	})

	It("does part A", func() {
		running := make(chan struct{})
		result := make(chan int, 1)

		go func() {
			r.RunProg(prog)
			close(running)
		}()

		go func() {
			for {
				open := false
				select {
				case <-running:
				default:
					open = true
				}
				if !open {
					result <- r.Visited()
					break
				}

				r.Move()
			}
		}()

		res := <-result
		Expect(res).To(Equal(2219))
	})

	It("does part B", func() {
		running := make(chan struct{})
		result := make(chan struct{})

		r.Set(grid.NewCoord(0, 0), robot.White)

		go func() {
			r.RunProg(prog)
			close(running)
		}()

		go func() {
			for {
				open := false
				select {
				case <-running:
				default:
					open = true
				}
				if !open {
					close(result)
					break
				}

				r.Move()
			}
		}()

		<-result
		output := r.GridToString()
		expectedOutput := ` #  #  ##  #### #  # #     ##  ###  ####   
 #  # #  # #    #  # #    #  # #  # #      
 #### #  # ###  #  # #    #  # #  # ###    
 #  # #### #    #  # #    #### ###  #      
 #  # #  # #    #  # #    #  # #    #      
 #  # #  # #     ##  #### #  # #    ####   
`
		Expect(output).To(Equal(expectedOutput))
	})
})
