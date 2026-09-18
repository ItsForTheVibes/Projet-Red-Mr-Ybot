package utils

import (
	"Projet-Red/src/character"
	"fmt"
)

func DisplayInfo(c character.Character) {
	fmt.Println()
	Header("PROJECT RED // OPERATOR STATUS")
	fmt.Println()

	fmt.Println(Cyan + "IDENTITY" + Reset)
	fmt.Println("├─ Alias       :", c.Name)
	fmt.Println("├─ Affiliation :", c.Class)
	fmt.Println("└─ Level       :", c.Level)

	fmt.Println()
	fmt.Println(Cyan + "SYSTEM" + Reset)
	fmt.Printf("├─ Integrity   : %s %d/%d\n",
		HealthBar(c.CurrentHP, c.MaxHP),
		c.CurrentHP,
		c.MaxHP,
	)
	fmt.Println("├─ Ethereum    :", c.Ethereum, "ETH")
	fmt.Println("└─ Inventory   :", len(c.Inventory), "items")

	fmt.Println()
	fmt.Println(Cyan + "PAYLOADS" + Reset)

	if len(c.Skill) == 0 {
		fmt.Println("└─ None")
	} else {
		for i, skill := range c.Skill {
			if i == len(c.Skill)-1 {
				fmt.Println("└─", skill)
			} else {
				fmt.Println("├─", skill)
			}
		}
	}

	fmt.Println()
}

func Menu(c character.Character) {
	for {
		var choix int

		fmt.Println()
		fmt.Println("===== MENU =====")
		fmt.Println()
		fmt.Println("1 - Display the characters information")
		fmt.Println("2 - Access the inventory")
		fmt.Println("3 - Quit")
		fmt.Println()
		fmt.Print("Choice : ")

		fmt.Scanln(&choix)

		switch choix {
		case 1:
			DisplayInfo(c)

		case 2:
			fmt.Println("===== Inventory =====")
			fmt.Println(c.Inventory)

		case 3:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
