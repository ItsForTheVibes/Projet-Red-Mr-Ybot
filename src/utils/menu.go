package utils

import (
	"Projet-Red/src/character"
	"fmt"
)

func DisplayInfo(c character.Character) {
	fmt.Println("==== Character Info ====")
	fmt.Println()

	fmt.Println("Name:", c.Name)
	fmt.Println("Class:", c.Class)
	fmt.Println("Level:", c.Level)
	fmt.Println("HP:", c.CurrentHP, "/", c.MaxHP)
	fmt.Println("Inventory:", c.Inventory)
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
