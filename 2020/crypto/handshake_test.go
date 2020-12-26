package crypto_test

import (
	"io"
	"strings"

	"github.com/kieron-dev/adventofcode/2020/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handshake", func() {
	var (
		data      io.Reader
		handshake crypto.Handshake
	)

	BeforeEach(func() {
		data = strings.NewReader(`
5764801
17807724
`)

		handshake = crypto.NewHandshake()
		handshake.Load(data)
	})

	It("can get the card loop no.", func() {
		Expect(handshake.CardLoopNum()).To(Equal(8))
	})

	It("can get the room loop no.", func() {
		Expect(handshake.RoomLoopNum()).To(Equal(11))
	})

	It("can get the encrpytion key", func() {
		Expect(handshake.EncryptionKey()).To(Equal(14897079))
	})
})
