package utils

import (
	"Projet-Red/src/character"
	"fmt"
)

func DisplayInfo(c character.Character) {
	color := ClassColor(c.Class)

	fmt.Println()
	ClassHeader("PROJECT RED // COMMAND CENTER", c.Class)
	fmt.Println()

	fmt.Println(color + "IDENTITY" + Reset)
	fmt.Println("├─ Alias       :", c.Name)
	fmt.Println("├─ Affiliation :", c.Class)
	fmt.Println("└─ Level       :", c.Level)

	fmt.Println()
	fmt.Println(color + "SYSTEM" + Reset)
	fmt.Printf("├─ Integrity   : %s %d/%d\n",
		HealthBar(c.CurrentHP, c.MaxHP),
		c.CurrentHP,
		c.MaxHP,
	)
	fmt.Println("├─ Ethereum    :", c.Ethereum, "ETH")
	fmt.Println("└─ Inventory   :", len(c.Inventory), "items")

	fmt.Println()
	fmt.Println(color + "PAYLOADS" + Reset)

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

		ClearScreen()

		Header("PROJECT RED // COMMAND CENTER")
		fmt.Println()

		fmt.Println("  [1] Operator Status")
		fmt.Println("  [2] Inventory")
		fmt.Println("  [3] Disconnect")
		fmt.Println()

		Prompt()
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			ClearScreen()
			DisplayInfo(c)

			fmt.Println()
			fmt.Println("Press ENTER to return...")
			fmt.Scanln()

		case 2:
			ClearScreen()

			Header("STORAGE // INVENTORY")
			fmt.Println()

			if len(c.Inventory) == 0 {
				fmt.Println("Inventory empty.")
			} else {
				for i, item := range c.Inventory {
					fmt.Printf("  [%02d] %s\n", i+1, item)
				}
			}

			fmt.Println()
			fmt.Println("Press ENTER to return...")
			fmt.Scanln()

		case 3:
			ClearScreen()
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
