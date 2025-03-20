package main

import (
	"fmt"
	"go-desing-pattern/factory_method/gun"
)

func main() {
	ak47, _ := gun.GetGun("ak47")
	musket, _ := gun.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun gun.IGun) {
	fmt.Printf("Gun: %s", gun.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", gun.GetPower())
	fmt.Println()
}
