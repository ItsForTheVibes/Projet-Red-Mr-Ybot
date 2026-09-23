package utils

import (
	"Projet-Red/src/character"
	"Projet-Red/src/combat"
	"Projet-Red/src/economy"
	"fmt"
)

func DisplayInfo(c *character.Character) {
	color := ClassColor(c.Class)

	fmt.Println()

	Header("PROJECT RED // OPERATOR STATUS")

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

	fmt.Println(
		"├─ Ethereum    :",
		c.Ethereum,
		"ETH",
	)

	fmt.Printf(
		"└─ Storage     : %d/%d\n",
		len(c.Inventory),
		c.InventoryCapacity,
	)

	fmt.Println()

	fmt.Println(color + "EQUIPMENT" + Reset)

	fmt.Println(
		"├─ Head  :",
		displayEmpty(c.Equipment.Head),
	)

	fmt.Println(
		"├─ Chest :",
		displayEmpty(c.Equipment.Chest),
	)

	fmt.Println(
		"└─ Feet  :",
		displayEmpty(c.Equipment.Feet),
	)

	fmt.Println()

	fmt.Println(color + "EXPLOITS" + Reset)

	for i, skill := range c.Skill {
		if i == len(c.Skill)-1 {
			fmt.Println("└─", skill)
		} else {
			fmt.Println("├─", skill)
		}
	}
}

func Menu(c *character.Character) {
	for {
		ClearScreen()

		Header("PROJECT RED // COMMAND CENTER")

		fmt.Println()
		fmt.Printf(
			"Operator: %s // %s\n",
			c.Name,
			c.Class,
		)

		fmt.Printf(
			"HP: %d/%d // ETH: %d\n",
			c.CurrentHP,
			c.MaxHP,
			c.Ethereum,
		)

		fmt.Println()

		fmt.Println("[1] Operator Status")
		fmt.Println("[2] Storage / Inventory")
		fmt.Println("[3] Darknet Market")
		fmt.Println("[4] Hardware Lab")
		fmt.Println("[5] Training Network")
		fmt.Println("[6] Disconnect")

		fmt.Println()

		var choice int

		Prompt()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ClearScreen()

			DisplayInfo(c)

			WaitForEnter()

		case 2:
			ClearScreen()

			character.AccessInventory(c)

			WaitForEnter()

		case 3:
			ClearScreen()

			economy.Merchant(c)

		case 4:
			ClearScreen()

			economy.Blacksmith(c)

		case 5:
			ClearScreen()

			combat.TrainingFight(c)

			WaitForEnter()

		case 6:
			ClearScreen()

			fmt.Println("[+] Connection terminated.")
			return

		default:
			Error("Invalid command.")

			WaitForEnter()
		}
	}
}

func displayEmpty(value string) string {
	if value == "" {
		return "[EMPTY]"
	}

	return value
}