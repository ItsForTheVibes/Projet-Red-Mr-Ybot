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
	fmt.Println("When consumed, receive 10 HP")
	fmt.Println("2 - Corruption Script")
	fmt.Println("Description:")
	fmt.Println("When consumed, Enemy receives 10 poison damage")
	fmt.Println("3 - Back")
	fmt.Print("Choice : ")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		c.Inventory = append(c.Inventory, "AntiVirus")
		fmt.Println("You received : AntiVirus")

	case 2:
		c.Inventory = append(c.Inventory, "Corruption Script")
		fmt.Println("You received : Corruption Script")

	case 3:
		return

	default:
		fmt.Println("Invalid choice.")
	}
}
