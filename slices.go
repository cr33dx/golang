package main

import "golang.org/x/tour/pic"


func main() {
    pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i,_ := range slice{
		slice[i] = make([]uint8,dx)
		for y,_ := range slice[i]{
			slice[i][y] = (uint8)((i+y)*(i-y))
		}
	}
	return slice
}
