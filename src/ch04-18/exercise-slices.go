package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var row []uint8
		for j := 0; j < dx; j++ {
			pixel := j % 4
			row = append(row, uint8(pixel))
		}
		pic = append(pic, row)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
