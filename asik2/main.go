package main

import (
	"fmt"

	abstractfactory "asik2/abstract_factory"
	"asik2/factory"
)

func runFactoryMethod() {
	fmt.Println("=== Factory Method ===")

	branches := []factory.IRentalBranch{
		factory.NewEconomyBranch("Airport — Economy Desk"),
		factory.NewSuvBranch("Downtown — SUV Specialists"),
		factory.NewVanBranch("Logistics Hub — Vans"),
	}

	contracts := []factory.RentalContract{
		factory.Rent(branches[0], 3),
		factory.Rent(branches[1], 2),
		factory.Rent(branches[2], 5),
	}

	for i, c := range contracts {
		fmt.Printf("Contract %d:\n", i+1)
		fmt.Println(c.String())
		fmt.Println()
	}
}

func runAbstractFactory() {
	fmt.Println("=== Abstract Factory ===")

	economy := abstractfactory.NewEconomyPackageFactory("Economy Package")
	luxury := abstractfactory.NewLuxuryPackageFactory("Luxury Package")
	family := abstractfactory.NewFamilySuvPackageFactory("Family SUV Package")

	agreements := []abstractfactory.RentalAgreement{
		abstractfactory.BuildAgreement(economy, 3),
		abstractfactory.BuildAgreement(luxury, 2),
		abstractfactory.BuildAgreement(family, 5),
	}

	for i, a := range agreements {
		fmt.Printf("Agreement %d:\n", i+1)
		fmt.Println(a.String())
		fmt.Println()
	}
}

func main() {
	runFactoryMethod()
	fmt.Println("-----------------------------------")
	runAbstractFactory()
}
