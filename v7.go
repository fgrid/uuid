package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// NewV4 creates a new UUID with variant 4 as described in RFC 4122. Variant 4 based on pure random bytes.
func NewV7() *UUID {
	time_buf := make([]byte, 8)
	now := time.Now()
	binary.BigEndian.PutUint64(time_buf[:], uint64(now.UnixMilli()))

	rand_buf := make([]byte, 10)
	rand.Read(rand_buf)

	var uuid UUID
	copy(uuid[:6], time_buf[2:])
	copy(uuid[6:], rand_buf[:])
	uuid[6] = (uuid[6] & 0x0f) | 0x70
	uuid.variantRFC9562()
	return &uuid
}
