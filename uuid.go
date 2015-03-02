package uuid

import (
	"encoding/binary"
	"fmt"
)

type UUID [16]byte

func (u *UUID) Version() int {
	return int(binary.BigEndian.Uint16(u[6:8]) >> 12)
}

func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func (u *UUID) variantRFC1422() {
	u[8] = (u[8] & 0x3f) | 0x80
}
