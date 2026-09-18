package economy

import (
	"Projet-Red/src/character"
	"fmt"
)

func Merchant(c *character.Character) {
	var choix int

	fmt.Println("===== MARCHAND =====")
	fmt.Println("1 - Potion de vie - Gratuit")
	fmt.Println("2 - Retour")
	fmt.Print("Choix : ")

	fmt.Scanln(&choix)

	switch choix {
	case 1:
		c.Inventory = append(c.Inventory, "Potion de vie")
		fmt.Println("Vous avez obtenu : Potion de vie")

	case 2:
		return

	default:
		fmt.Println("Choix invalide.")
	}
}
