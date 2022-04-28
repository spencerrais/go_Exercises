package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	dx_data := make([]uint8, dx)
	dy_data := make([][]uint8, dy)

	for i := range dy_data {
		for j := range dx_data {
			dx_data[j] = uint8((j + i) / 2)
		}
		dy_data[i] = dx_data
	}
	fmt.Println("%v\n", dy_data)
	return dy_data
}

func main() {
	pic.Show(Pic)
}
