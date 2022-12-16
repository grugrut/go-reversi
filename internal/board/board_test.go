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

func TestRotateHorizontal(t *testing.T) {
	var b, got, want uint64
	b = 0x1020408001020408
	got = RotateHorizontal(b)
	want = 0x0804020180402010

	if got != want {
		t.Errorf("RotateHorizontal(%b) == %b, want %b", b, got, want)
	}
}

func TestGetPlacable(t *testing.T) {
	var self, oppo, got, want uint64

	// 1つ返せるパターン
	self = 0x0000200000000000
	oppo = 0x0070507000000000
	want = 0xA8008800A8000000

	got = GetPlacable(self, oppo)
	if got != want {
		t.Errorf("GetPlacable(%b, %b) == %b, want %b", self, oppo, got, want)
	}

	// 2つ返せるパターン
	self = 0x0000001000000000
	oppo = 0x007C7C6C7C7C0000
	want = 0x9200008200009200

	got = GetPlacable(self, oppo)
	if got != want {
		t.Errorf("GetPlacable(%b, %b) == %b, want %b", self, oppo, got, want)
	}

	// 返せる位置には自石があり返せないパターン
	self = 0x007C4454447C0000
	oppo = 0x0000382838000000
	want = 0x0000000000000000

	got = GetPlacable(self, oppo)
	if got != want {
		t.Errorf("GetPlacable(%b, %b) == %b, want %b", self, oppo, got, want)
	}

	// 辺や角で返せる位置が無いパターン
	self = 0x004A02400002CA00
	oppo = 0x89008180810100FB
	want = 0x0000000000000000

	got = GetPlacable(self, oppo)
	if got != want {
		t.Errorf("GetPlacable(%b, %b) == %b, want %b", self, oppo, got, want)
	}
}
