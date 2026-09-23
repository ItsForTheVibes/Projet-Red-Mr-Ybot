package economy

import (
	"Projet-Red/src/character"
	"fmt"
)

func Merchant(c *character.Character) {
	for {
		var choice int

		fmt.Println()
		fmt.Println("===== MERCHANT =====")
		fmt.Println()
		fmt.Println("Ethereum :", c.Ethereum)
		fmt.Println()

		fmt.Println("1 - AntiVirus - Free")
		fmt.Println("2 - Corruption Script - Free")
		fmt.Println("3 - Exploit Script - 25 Ethereum")
		fmt.Println("4 - Encrypted Data - 4 Ethereum")
		fmt.Println("5 - Firewall Module - 7 Ethereum")
		fmt.Println("6 - Security Token - 3 Ethereum")
		fmt.Println("7 - API Key - 1 Ethereum")
		fmt.Println("8 - Livre de Sort : Boule de Feu - 10 Ethereum")
		fmt.Println("9 - Back")
		fmt.Println()

		fmt.Print("Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			buyItem(c, "AntiVirus", 0)

		case 2:
			buyItem(c, "Corruption Script", 0)

		case 3:
			buyItem(c, "Exploit Script", 25)

		case 4:
			buyItem(c, "Encrypted Data", 4)

		case 5:
			buyItem(c, "Firewall Module", 7)

		case 6:
			buyItem(c, "Security Token", 3)

		case 7:
			buyItem(c, "API Key", 1)

		case 8:
			buyItem(
				c,
				"Livre de Sort : Boule de Feu",
				10,
			)

		case 9:
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func buyItem(c *character.Character, item string, price int) {
	if c.Ethereum < price {
		fmt.Println("Not enough Ethereum.")
		return
	}

	if len(c.Inventory) >= 10 {
		fmt.Println("Inventory full.")
		return
	}

	c.Ethereum -= price

	character.AddItem(c, item)

	if price > 0 {
		fmt.Printf(
			"You bought %s for %d Ethereum.\n",
			item,
			price,
		)
	} else {
		fmt.Printf(
			"You received %s.\n",
			item,
		)
	}
}