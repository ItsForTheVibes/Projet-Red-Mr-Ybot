package utils

import (
	"Projet-Red/src/character"
	"Projet-Red/src/combat"
	"Projet-Red/src/economy"
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

	fmt.Printf(
		"├─ Integrity   : %s %d/%d\n",
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

func Menu(c *character.Character) {
	for {
		var choix int

		ClearScreen()

		Header("PROJECT RED // COMMAND CENTER")
		fmt.Println()

		fmt.Println("  [1] Operator Status")
		fmt.Println("  [2] Inventory")
		fmt.Println("  [3] Merchant")
		fmt.Println("  [4] Blacksmith")
		fmt.Println("  [5] Combat")
		fmt.Println("  [6] Disconnect")
		fmt.Println()

		Prompt()
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			ClearScreen()

			DisplayInfo(*c)

			waitForEnter()

		case 2:
			ClearScreen()

			character.AccessInventory(c)

			waitForEnter()

		case 3:
			ClearScreen()

			economy.Merchant(c)

		case 4:
			ClearScreen()

			economy.Blacksmith(c)

		case 5:
			ClearScreen()

			combat.StartCombat(c)

			waitForEnter()

		case 6:
			ClearScreen()

			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
			waitForEnter()
		}
	}
}

func waitForEnter() {
	fmt.Println()
	fmt.Println("Press ENTER to return...")
	fmt.Scanln()
}