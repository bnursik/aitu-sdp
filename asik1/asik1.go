package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Color string
	Year  int
}

type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetModel(model string) CarBuilder
	SetColor(color string) CarBuilder
	SetYear(year int) CarBuilder
	build() Car
}

type ConcreteCarBuilder struct {
	car Car
}

func (ConcreteCarBuilder *ConcreteCarBuilder) SetBrand(brand string) CarBuilder {
	ConcreteCarBuilder.car.Brand = brand
	return ConcreteCarBuilder
}

func (ConcreteCarBuilder *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	ConcreteCarBuilder.car.Model = model
	return ConcreteCarBuilder
}

func (ConcreteCarBuilder *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	ConcreteCarBuilder.car.Color = color
	return ConcreteCarBuilder
}

func (ConcreteCarBuilder *ConcreteCarBuilder) SetYear(year int) CarBuilder {
	ConcreteCarBuilder.car.Year = year
	return ConcreteCarBuilder
}

func (ConcreteCarBuilder *ConcreteCarBuilder) build() Car {
	return ConcreteCarBuilder.car
}

type CarDirector struct{}

func (d *CarDirector) ConstructSportsCar_Porsche911(builder CarBuilder) Car {
	return builder.SetBrand("Porsche").
		SetModel("911").
		SetColor("Red").
		SetYear(2021).
		build()
}

func main() {
	director := CarDirector{}
	builder := &ConcreteCarBuilder{}

	ConstructSportsCar_Porsche911 := director.ConstructSportsCar_Porsche911(builder)
	fmt.Println("Sports Car:", ConstructSportsCar_Porsche911)

	customCar_TeslaS := builder.SetBrand("Tesla").SetModel("Model S").SetColor("Black").SetYear(2022).build()
	fmt.Println("Custom Car:", customCar_TeslaS)
}
