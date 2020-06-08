package main

import (
	"GoCourse/homework-3/answer1-2/models"
	"fmt"
	"reflect"
)

func main() {
	car := models.Car{
		Base:      models.Base{Brand: "Volvo", Year: 1980, BootVolume: 30.5, BootFill: 20.8, IsWindowOpend: true},
		ClassName: "Coupe"}
	truck := models.Truck{
		Base:       models.Base{Brand: "Bmw", Year: 1990, BootVolume: 100, BootFill: 90, IsEngineStarted: true},
		WheelCount: 14}
	fmt.Println(car)
	fmt.Println(truck)
	car.Base.IsEngineStarted = true
	truck.Base.IsWindowOpend = false
	fmt.Println(car, reflect.TypeOf(car))
	fmt.Println(truck, reflect.TypeOf(truck))
}
