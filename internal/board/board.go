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
