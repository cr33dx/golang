package main

import "fmt"

type vehicle struct {
	name string
	tyre int
	breaks string
}

func (v vehicle) iam() string{
	return(fmt.Sprintf("hi this is %s i have %d tyres", v.name, v.tyre))
}

func main(){
	car := vehicle{name:"swift",tyre:4,breaks:"yes"}
	fmt.Println(car.iam())
}
