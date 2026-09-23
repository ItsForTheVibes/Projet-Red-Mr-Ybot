package character

import "fmt"

func AccessInventory(c *Character) {
	fmt.Println("=== INVENTAIRE ===")
	fmt.Println()

	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Println("0. Retour")
	fmt.Println()

	var choice int

	fmt.Print("Choisissez un objet à utiliser : ")
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(c.Inventory) {
		fmt.Println("Choix invalide.")
		return
	}

	item := c.Inventory[choice-1]

	switch item {
	case "Chapeau de l'aventurier",
		"Tunique de l'aventurier",
		"Bottes de l'aventurier":

		c.EquipItem(item)
		return

	case "Livre de Sort : Boule de Feu":

		if spellBook(c) {
			c.Inventory = append(
				c.Inventory[:choice-1],
				c.Inventory[choice:]...,
			)
		}

		return

	default:
		fmt.Println("Cet objet ne peut pas être utilisé ici.")
	}
}

func spellBook(c *Character) bool {
	for _, skill := range c.Skill {
		if skill == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de Feu.")
			return false
		}
	}

	c.Skill = append(c.Skill, "Boule de Feu")

	fmt.Println("Vous avez appris le sort Boule de Feu !")

	return true
}

func AddItem(c *Character, item string) {
	if len(c.Inventory) >= 10 {
		fmt.Printf(
			"Inventaire plein ! Impossible d'ajouter '%s'.\n",
			item,
		)

		return
	}

	c.Inventory = append(c.Inventory, item)

	fmt.Printf("%s ajouté à l'inventaire.\n", item)
}