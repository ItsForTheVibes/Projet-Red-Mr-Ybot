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

	fmt.Println("3 - Exploit Script - 25 Ethereum")
	fmt.Println("    A powerful offensive script.")
	fmt.Println()

	fmt.Println("4 - Encrypted Data - 4 Ethereum")
	fmt.Println("    Encrypted information used for crafting.")
	fmt.Println()

	fmt.Println("5 - Firewall Module - 7 Ethereum")
	fmt.Println("    A security component used for crafting.")
	fmt.Println()

	fmt.Println("6 - Security Token - 3 Ethereum")
	fmt.Println("    An authentication token used for crafting.")
	fmt.Println()

	fmt.Println("7 - API Key - 1 Ethereum")
	fmt.Println("    An access key used for crafting.")
	fmt.Println()

	fmt.Println("8 - Livre de Sort : Boule de Feu - 10 Ethereum")
	fmt.Println("    Permet d'apprendre le sort Boule de Feu.")
	fmt.Println()

	fmt.Println("9 - Back")
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
		c.Inventory = append(c.Inventory, "Exploit Script")
		fmt.Println("You received: Exploit Script")

	case 4:
		c.Inventory = append(c.Inventory, "Encrypted Data")
		fmt.Println("You received: Encrypted Data")

	case 5:
		c.Inventory = append(c.Inventory, "Firewall Module")
		fmt.Println("You received: Firewall Module")

	case 6:
		c.Inventory = append(c.Inventory, "Security Token")
		fmt.Println("You received: Security Token")

	case 7:
		c.Inventory = append(c.Inventory, "API Key")
		fmt.Println("You received: API Key")

	case 8:
		c.Inventory = append(c.Inventory, "Livre de Sort : Boule de Feu")
		fmt.Println("You received: Livre de Sort : Boule de Feu")
		
	case 9:
		return

	default:
		fmt.Println("Invalid choice.")
	}
}
