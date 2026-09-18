package utils

import (
	"fmt"

	"Projet-Red/src/character"
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
