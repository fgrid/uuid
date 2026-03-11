package uuid

import (
	"strings"
	"testing"
	"time"
)

func TestNewV7(t *testing.T) {
	var old *UUID
	for i := range 10 {
		time.Sleep(2 * time.Millisecond)
		uuid := NewV7()
		if uuid.Version() != 7 {
			t.Errorf("invalid version %d - expected 7", uuid.Version())
		}
		if i > 0 && strings.Compare(old.String(), uuid.String()) != -1 {
			t.Errorf("v7 uuid not increasing as expected %q -> %q", old.String(), uuid.String())
		}
		old = uuid
	}
}

func BenchmarkNewV7(b *testing.B) {
	for b.Loop() {
		NewV7()
	}
}
