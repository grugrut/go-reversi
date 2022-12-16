package board

import (
	"fmt"
)

/*
    A   B   C   D   E   F   G   H
  ┌───┬───┬───┬───┬───┬───┬───┬───┐
1 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
2 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
3 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
4 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
5 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
6 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
7 │   │   │   │   │   │   │   │   │
  ├───┼───┼───┼───┼───┼───┼───┼───┤
8 │   │   │   │   │   │   │   │   │
  └───┴───┴───┴───┴───┴───┴───┴───┘
*/

var black uint64
var white uint64

const FIRST_CELL uint64 = 0x8000000000000000

// Init は盤面の初期化処理
func Init() {
	black = 0x0810000000
	white = 0x1008000000
}

// PrintBoard は盤面の描画をする
func PrintBoard() {
	fmt.Println("    A   B   C   D   E   F   G   H  ")
	fmt.Println("  ╋───╋───╋───╋───╋───╋───╋───╋───╋")
	for i := 0; i < 8; i++ {
		fmt.Print(i+1, " ")
		for j := 0; j < 8; j++ {
			fmt.Print("┃ ")
			if black&(FIRST_CELL>>(i*8+j)) != 0 {
				fmt.Print("●")
			} else if white&(FIRST_CELL>>(i*8+j)) != 0 {
				fmt.Print("○")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println("┃")
		fmt.Println("  ╋───╋───╋───╋───╋───╋───╋───╋───╋")
	}
}

// Count はビットの数を返す
func Count(b uint64) int {
	b = b - ((b >> 1) & 0x5555555555555555)
	b = (b & 0x3333333333333333) + ((b >> 2) & 0x3333333333333333)
	b = (b + (b >> 4)) & 0x0f0f0f0f0f0f0f0f
	b += (b >> 8)
	b += (b >> 16)
	b += (b >> 32)

	return (int)(b & 0x7f)
}

// Rotate90は盤面を反時計回りに90度回転する
func Rotate90(b uint64) uint64 {
	var tmp uint64
	var k1 uint64 = 0x5500550055005500
	var k2 uint64 = 0x3333000033330000
	var k3 uint64 = 0x0f0f0f0f00000000

	tmp = k3 & (b ^ (b << 28))
	b ^= tmp ^ (tmp >> 28)
	tmp = k2 & (b ^ (b << 14))
	b ^= tmp ^ (tmp >> 14)
	tmp = k1 & (b ^ (b << 7))
	b ^= tmp ^ (tmp >> 7)

	var l1 uint64 = 0x00ff00ff00ff00ff
	var l2 uint64 = 0x0000ffff0000ffff

	b = ((b >> 8) & l1) | ((b & l1) << 8)
	b = ((b >> 16) & l2) | ((b & l2) << 16)
	b = (b >> 32) | (b << 32)

	return b
}

// RotateHorizontal は盤面を180度反転させる
func RotateHorizontal(b uint64) uint64 {
	var k1 uint64 = 0x5555555555555555
	var k2 uint64 = 0x3333333333333333
	var k3 uint64 = 0x0f0f0f0f0f0f0f0f

	b = ((b >> 1) & k1) + 2*(b&k1)
	b = ((b >> 2) & k2) + 4*(b&k2)
	b = ((b >> 4) & k3) + 16*(b&k3)

	return b
}

// GetPlacable はself番のプレイヤーが配置できる箇所を返す
func GetPlacable(self uint64, oppo uint64) uint64 {
	var result, tmp, guard uint64

	// left
	guard = 0x7e7e7e7e7e7e7e7e
	tmp = (self << 1) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp << 1) & oppo & guard
	}
	result |= (tmp << 1) & ^(self | oppo)

	// up left
	guard = 0x007e7e7e7e7e7e7e
	tmp = (self << 9) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp << 9) & oppo & guard
	}
	result |= (tmp << 9) & ^(self | oppo)

	// up
	guard = 0x00ffffffffffffff
	tmp = (self << 8) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp << 8) & oppo & guard
	}
	result |= (tmp << 8) & ^(self | oppo)

	// up right
	guard = 0x007e7e7e7e7e7e7e
	tmp = (self << 7) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp << 7) & oppo & guard
	}
	result |= (tmp << 7) & ^(self | oppo)

	// right
	guard = 0x7e7e7e7e7e7e7e7e
	tmp = (self >> 1) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp >> 1) & oppo & guard
	}
	result |= (tmp >> 1) & ^(self | oppo)

	// down right
	guard = 0x7e7e7e7e7e7e7e00
	tmp = (self >> 9) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp >> 9) & oppo & guard
	}
	result |= (tmp >> 9) & ^(self | oppo)

	// down
	guard = 0xffffffffffffff00
	tmp = (self >> 8) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp >> 8) & oppo & guard
	}
	result |= (tmp >> 8) & ^(self | oppo)

	// down left
	guard = 0x7e7e7e7e7e7e7e00
	tmp = (self >> 7) & oppo & guard
	for i := 0; i < 5; i++ {
		tmp |= (tmp >> 7) & oppo & guard
	}
	result |= (tmp >> 7) & ^(self | oppo)

	return result
}
