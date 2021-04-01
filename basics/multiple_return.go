package main

import "fmt"

func test() (int, int, string) {
	return 12, 12, "testing"
}

func main(){
	x, y, z := test()
	fmt.Println(x, z, y)
}
