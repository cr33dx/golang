package main

import "fmt"

type car struct{
	max_speed, horse_power uint16
	color string
	hatchback bool
}

func (c car) carInfo(owner string) string{
	return(fmt.Sprintf("owner of car is %s max speed is %d color is %s & hatchback is %d", owner, c.max_speed, c.color, c.hatchback))
}

func main(){
	a_car := car{max_speed: 250, color: "red", horse_power: 220, hatchback: true}
	fmt.Println(a_car.carInfo("sameer"))
}
