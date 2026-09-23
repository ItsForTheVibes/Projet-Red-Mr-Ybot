package economy

import (
	"Projet-Red/src/character"
	"fmt"
)

func Blacksmith(c *character.Character) {
	for {
		var choice int

		fmt.Println()
		fmt.Println("===== BLACKSMITH =====")
		fmt.Println()
		fmt.Println("Ethereum :", c.Ethereum)
		fmt.Println()

		fmt.Println("PAYLOADS")
		fmt.Println("1 - AntiVirus")
		fmt.Println("2 - Corruption Script")
		fmt.Println("3 - Exploit Script")
		fmt.Println("4 - Encrypted Data")
		fmt.Println("5 - Firewall Module")
		fmt.Println("6 - Security Token")
		fmt.Println("7 - API Key")
		fmt.Println("8 - Livre de Sort : Boule de Feu")
		fmt.Println()

		fmt.Println("EQUIPMENT")
		fmt.Println("9 - Chapeau de l'aventurier")
		fmt.Println("10 - Tunique de l'aventurier")
		fmt.Println("11 - Bottes de l'aventurier")
		fmt.Println()

		fmt.Println("12 - Back")
		fmt.Println()

		fmt.Println("Crafting cost: 5 Ethereum")
		fmt.Print("Choice: ")

		fmt.Scanln(&choice)

		var item string

		switch choice {
		case 1:
			item = "AntiVirus"

		case 2:
			item = "Corruption Script"

		case 3:
			item = "Exploit Script"

		case 4:
			item = "Encrypted Data"

		case 5:
			item = "Firewall Module"

		case 6:
			item = "Security Token"

		case 7:
			item = "API Key"

		case 8:
			item = "Livre de Sort : Boule de Feu"

		case 9:
			item = "Chapeau de l'aventurier"

		case 10:
			item = "Tunique de l'aventurier"

		case 11:
			item = "Bottes de l'aventurier"

		case 12:
			return

		default:
			fmt.Println("Invalid choice.")
			continue
		}

		craftItem(c, item)
	}
}

func craftItem(c *character.Character, item string) {
	const craftPrice = 5

	if c.Ethereum < craftPrice {
		fmt.Println("Not enough Ethereum.")
		return
	}

	if len(c.Inventory) >= 10 {
		fmt.Println("Inventory full.")
		return
	}

	c.Ethereum -= craftPrice

	character.AddItem(c, item)

	fmt.Printf(
		"%s crafted successfully!\n",
		item,
	)

	fmt.Printf(
		"Cost: %d Ethereum\n",
		craftPrice,
	)
}