// Package crypto cracks a hotel door
package crypto

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

const bigPrime = 20201227

type Handshake struct {
	cardPub  int
	roomPub  int
	cardLoop int
	roomLoop int
}

func NewHandshake() Handshake {
	return Handshake{}
}

func (h *Handshake) Load(data io.Reader) {
	scanner := bufio.NewScanner(data)
	numList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("atoi failed: %v", err)
		}

		numList = append(numList, num)
	}

	if len(numList) != 2 {
		log.Fatalf("expected len 2: %v", numList)
	}

	h.cardPub = numList[0]
	h.roomPub = numList[1]
}

func transform(loop, subject int) int {
	n := 1
	for i := 0; i < loop; i++ {
		n *= subject
		n %= bigPrime
	}

	return n
}

func crack(pub int) int {
	loop := 1
	n := 1

	for {
		n = (n * 7) % bigPrime
		if n == pub {
			break
		}
		loop++
	}

	return loop
}

func (h *Handshake) CardLoopNum() int {
	if h.cardLoop == 0 {
		h.cardLoop = crack(h.cardPub)
	}

	return h.cardLoop
}

func (h *Handshake) RoomLoopNum() int {
	if h.roomLoop == 0 {
		h.roomLoop = crack(h.roomPub)
	}

	return h.roomLoop
}

func (h *Handshake) EncryptionKey() int {
	return transform(h.RoomLoopNum(), h.cardPub)
}
