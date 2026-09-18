package economy

import (
	"Projet-Red/src/character"
	"fmt"
)

func Merchant(c *character.Character) {
	var choice int

	fmt.Println("===== MERCHANT =====")
	fmt.Println()

	fmt.Println("1 - AntiVirus - Free")
	fmt.Println("    Restores 10 HP when consumed.")
	fmt.Println()

	fmt.Println("2 - Corruption Script - Free")
	fmt.Println("    Inflicts 10 poison damage on an enemy.")
	fmt.Println()

	fmt.Println("3 - Back")
	fmt.Println()

	fmt.Print("Choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Inventory = append(c.Inventory, "AntiVirus")
		fmt.Println("You received: AntiVirus")

	case 2:
		c.Inventory = append(c.Inventory, "Corruption Script")
		fmt.Println("You received: Corruption Script")

	case 3:
		return

	default:
		fmt.Println("Invalid choice.")
	}
}
