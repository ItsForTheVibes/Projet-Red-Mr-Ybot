package character


import "fmt"

// Fonction qui affiche tous les objets de l'inventaire (Tâche 4)
func AccessInventory(c Character) {
	fmt.Println("=== INVENTAIRE ===")

	// Vérification si l'inventaire est vide
	if len(c.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
		return
	}

	// Parcours et affichage de chaque objet de la liste
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
