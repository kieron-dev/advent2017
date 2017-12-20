package assembly_test

import (
	"github.com/kieron-pivotal/advent2017/18/assembly"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assembly", func() {
	Context("single machine", func() {
		var (
			m *assembly.Machine
		)

		BeforeEach(func() {
			m = assembly.NewMachine(0)
		})

		Context("basic operations", func() {
			It("can set a register", func() {
				m.Execute("set a 1")
				Expect(m.GetRegister('a')).To(Equal(1))
			})

			It("can add to a register", func() {
				m.Execute("set a 1")
				m.Execute("add a 2")
				Expect(m.GetRegister('a')).To(Equal(3))
			})

			It("can multiply a register", func() {
				m.Execute("set a 2")
				m.Execute("mul a 3")
				Expect(m.GetRegister('a')).To(Equal(6))
			})

			It("can do modulo on a register", func() {
				m.Execute("set a 13")
				m.Execute("mod a 7")
				Expect(m.GetRegister('a')).To(Equal(6))
			})

			It("can handle a jgz", func() {
				Expect(func() {
					m.Execute("jgz 0 2")
				}).ShouldNot(Panic())
			})

			It("can set a register to the value of another", func() {
				m.Execute("set a 2")
				m.Execute("set b a")
				Expect(m.GetRegister('b')).To(Equal(2))
			})
		})

		Context("sets of ops", func() {
			It("can execute a sequence of ops", func() {
				m.AppendInstruction("set a 40")
				m.AppendInstruction("add a 2")
				m.Run()
				Expect(m.GetRegister('a')).To(Equal(42))
			})

			It("can jump", func() {
				m.AppendInstruction("set a 40")
				m.AppendInstruction("jgz 2 2")
				m.AppendInstruction("add a 2")
				m.Run()
				Expect(m.GetRegister('a')).To(Equal(40))
			})

			It("can skip a jump", func() {
				m.AppendInstruction("set a 40")
				m.AppendInstruction("jgz b 2")
				m.AppendInstruction("add a 2")
				m.Run()
				Expect(m.GetRegister('a')).To(Equal(42))
			})
		})

	})

	Context("sending and receiving", func() {

		var (
			m1 *assembly.Machine
			m2 *assembly.Machine
		)

		BeforeEach(func() {
			m1 = assembly.NewMachine(0)
			m2 = assembly.NewMachine(1)
			m1.Duet(m2)
		})

		It("sets up machines with correct IDs", func() {
			Expect(m1.GetRegister('p')).To(Equal(0))
			Expect(m2.GetRegister('p')).To(Equal(1))
		})

		It("will send from one machine to the other", func() {
			m1.Execute("snd 20")
			m2.Execute("rcv z")
			m2.Execute("snd p")
			m1.Execute("rcv b")
			Expect(m2.GetRegister('z')).To(Equal(20))
			Expect(m1.GetRegister('b')).To(Equal(1))
		})

		It("counts number of sends", func() {
			m1.Execute("snd 20")
			m2.Execute("rcv z")
			m2.Execute("snd p")
			m1.Execute("rcv b")
			Expect(m1.GetCount()).To(Equal(1))
		})

		It("correctly does example", func() {
			for _, s := range []string{
				"snd 1",
				"snd 2",
				"snd p",
				"rcv a",
				"rcv b",
				"rcv c",
				"rcv d",
			} {
				m1.AppendInstruction(s)
				m2.AppendInstruction(s)
			}

			go assembly.RunMachines([]*assembly.Machine{m1, m2})
			Eventually(m2.GetCount).Should(Equal(3))
		})

	})
})
