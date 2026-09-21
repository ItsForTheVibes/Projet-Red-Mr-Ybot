package character

import "fmt"
	

func AccessInventory(c *Character) {
	fmt.Println("=== INVENTAIRE ===")

	
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	var choice int
	print("Choisissez un objet à utiliser : ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(c.Inventory) {
		println("Choix invalide.")
		return
	}

	item := c.Inventory[choice-1]

	if item == "Livre de Sort : Boule de Feu" {
		spellBook(c)

		c.Inventory = append(c.Inventory[:choice-1], c.Inventory[choice:]...)
		return
	}

	fmt.Println("Cet objet ne peut pas être utilisé ici.")
}

func spellBook(character *Character) {
	for _, skill := range character.Skill {
		if skill == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de Feu.")
			return
		}
	}

	character.Skill = append(character.Skill, "Boule de Feu")
	println("Vous avez appris le sort Boule de Feu !")
}

func AddItem(c *Character, item string) {
	
	if len(c.Inventory) >= 10 {
		fmt.Printf(" Inventaire plein ! Impossible d'ajouter '%s'.\n", item)
		return
	}

	
	c.Inventory = append(c.Inventory, item)
	fmt.Printf(" %s ajouté à l'inventaire.\n", item)
}