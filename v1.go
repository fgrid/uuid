package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"time"
)

type stamp [10]byte

var (
	mac     []byte
	service chan chan stamp
)

const gregorianUnix = 122192928000000000 // nanoseconds between gregorion zero and unix zero

func init() {
	mac = make([]byte, 6)
	rand.Read(mac)
	service = make(chan chan stamp)
	go unique(service)
	i, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, d := range i {
		if len(d.HardwareAddr) == 6 {
			mac = d.HardwareAddr[:6]
			return
		}
	}
}

func NewV1() *UUID {
	var uuid UUID
	a := make(chan stamp)
	service <- a
	s := <-a
	copy(uuid[:4], s[4:])
	copy(uuid[4:6], s[2:4])
	copy(uuid[6:8], s[:2])
	uuid[6] = (uuid[6] & 0x0f) | 0x10
	copy(uuid[8:10], s[8:])
	copy(uuid[10:], mac)
	return &uuid
}

func unique(service chan chan stamp) {
	var (
		lastNanoTicks  uint64
		clock_sequence [2]byte
	)
	rand.Read(clock_sequence[:])

	for c := range service {
		var s stamp
		nanoTicks := uint64((time.Now().UTC().UnixNano() / 100) + gregorianUnix)
		if nanoTicks < lastNanoTicks {
			lastNanoTicks = nanoTicks
			rand.Read(clock_sequence[:])
		} else if nanoTicks == lastNanoTicks {
			lastNanoTicks = nanoTicks + 1
		} else {
			lastNanoTicks = nanoTicks
		}
		binary.BigEndian.PutUint64(s[:], lastNanoTicks)
		copy(s[8:], clock_sequence[:])
		c <- s
	}
}
