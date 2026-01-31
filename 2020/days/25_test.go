package days_test

import (
	"os"

	"github.com/kieron-dev/adventofcode/2020/crypto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("25", func() {
	var (
		data      *os.File
		handshake crypto.Handshake
	)

	BeforeEach(func() {
		var err error
		data, err = os.Open("./input25")
		Expect(err).NotTo(HaveOccurred())

		handshake = crypto.NewHandshake()
		handshake.Load(data)
	})

	AfterEach(func() {
		data.Close()
	})

	It("can do part A", func() {
		Expect(handshake.EncryptionKey()).To(Equal(12227206))
	})
})
