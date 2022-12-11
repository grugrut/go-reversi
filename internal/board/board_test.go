package board

import (
	"testing"
)

func TestCount(t *testing.T) {
	var b uint64
	b = 0x1000000000000000
	got := Count(b)
	want := 1

	if got != want {
		t.Errorf("Count(%v) == %d, want %d", b, got, want)
	}

	b = 0x1000000010000000
	got = Count(b)
	want = 2

	if got != want {
		t.Errorf("Count(%v) == %d, want %d", b, got, want)
	}

	b = 0x30000000F0000000
	got = Count(b)
	want = 6

	if got != want {
		t.Errorf("Count(%v) == %d, want %d", b, got, want)
	}

}

func TestRotate90(t *testing.T) {

}
