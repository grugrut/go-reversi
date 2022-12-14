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
		t.Errorf("Count(%b) == %d, want %d", b, got, want)
	}

}

func TestRotate90(t *testing.T) {
	var b, got, want uint64
	b = 0x1020400000000000
	got = Rotate90(b)
	want = 0x0000000080402000

	if got != want {
		t.Errorf("1 x Rotate90(%b) == %b, want %b", b, got, want)
	}

	got = Rotate90(Rotate90(b))
	want = 0x0000000000020408
	if got != want {
		t.Errorf("2 x Rotate90(%b) == %b, want %b", b, got, want)
	}

	got = Rotate90(Rotate90(Rotate90(b)))
	want = 0x0004020100000000
	if got != want {
		t.Errorf("3 x Rotate90(%b) == %b, want %b", b, got, want)
	}

	got = Rotate90(Rotate90(Rotate90(Rotate90(b))))
	want = 0x1020400000000000
	if got != want {
		t.Errorf("4 x Rotate90(%b) == %b, want %b", b, got, want)
	}
}
