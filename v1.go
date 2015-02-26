package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"time"
)

var (
	mac            []byte
	clock_sequence []byte
)

const gregorianUnix = 122192928000000000 // nanoseconds between gregorion zero and unix zero

func init() {
	mac = make([]byte, 6)
	clock_sequence = make([]byte, 2)
	rand.Read(mac)
	rand.Read(clock_sequence)
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
	stamp := (time.Now().UTC().UnixNano() / 100) + gregorianUnix
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(stamp))
	copy(uuid[:4], buf[4:])
	copy(uuid[4:6], buf[2:4])
	copy(uuid[6:8], buf[:2])
	uuid[6] = (uuid[6] & 0x0f) | 0x10
	copy(uuid[8:10], clock_sequence)
	copy(uuid[10:], mac)
	return &uuid
}
