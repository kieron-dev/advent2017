package assembly_test

import (
	"github.com/kieron-pivotal/advent2017/18/assembly"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assembly", func() {
	var (
		m *assembly.Machine
	)
	BeforeEach(func() {
		m = assembly.NewMachine()
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

		It("can play a sound", func() {
			Expect(func() { m.Execute("snd 12") }).ShouldNot(Panic())
		})

		It("rcv returns true on non-zero arg", func() {
			Expect(m.Execute("rcv 0")).To(BeFalse())
			Expect(m.Execute("rcv 1")).To(BeTrue())
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

	Context("sounds", func() {
		It("stores first rcv val with non-zero arg", func() {
			m.AppendInstruction("snd 40")
			m.AppendInstruction("rcv 0")
			m.AppendInstruction("snd 42")
			m.AppendInstruction("snd 39")
			m.AppendInstruction("rcv 1")
			m.Run()
			Expect(m.RecoverVal()).To(Equal(39))
		})
	})
})
