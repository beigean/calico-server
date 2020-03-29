// ref. https://go-tour-jp.appspot.com/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	canvas := make([][]uint8, dy, dy)
	for y := 0; y < dy; y++ {
		canvas[y] = make([]uint8, dx, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			canvas[x][y] = uint8((x + y) / 2)
			// canvas[x][y] = uint8(x * y)
			// canvas[x][y] = uint8(x ^ y)
		}
	}
	return canvas
}

func main() {
	pic.Show(Pic)
}
