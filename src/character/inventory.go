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
}


func AddItem(c *Character, item string) {
	
	if len(c.Inventory) >= 10 {
		fmt.Printf(" Inventaire plein ! Impossible d'ajouter '%s'.\n", item)
		return
	}

	
	c.Inventory = append(c.Inventory, item)
	fmt.Printf(" %s ajouté à l'inventaire.\n", item)
}