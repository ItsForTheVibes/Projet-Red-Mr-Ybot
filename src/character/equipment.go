package character

import (
	"Projet-Red/src/items"
	"fmt"
)

type Equipment struct {
	Head  string
	Chest string
	Feet  string
}

func EquipItem(c *Character, item string) {
	recipe, exists := items.FindRecipe(item)

	if !exists {
		fmt.Println("[-] This item cannot be equipped.")
		return
	}

	if CountItem(c, item) == 0 {
		fmt.Println("[-] Equipment not found in inventory.")
		return
	}

	var slot *string

	switch item {
	case items.NeuralVisor:
		slot = &c.Equipment.Head

	case items.FirewallJacket:
		slot = &c.Equipment.Chest

	case items.ProxyBoots:
		slot = &c.Equipment.Feet
	}

	if *slot != "" {
		oldItem := *slot
		oldRecipe, _ := items.FindRecipe(oldItem)

		c.MaxHP -= equipmentBonus(oldItem)

		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}

		AddItem(c, oldItem)

		fmt.Printf(
			"[>] %s returned to inventory.\n",
			oldItem,
		)

		_ = oldRecipe
	}

	RemoveItem(c, item, 1)

	*slot = item
	c.MaxHP += equipmentBonus(item)

	fmt.Printf("[+] Equipped: %s\n", item)
	fmt.Printf("[+] Maximum HP +%d\n", recipeHPBonus(recipe.Name))
}

func equipmentBonus(item string) int {
	return recipeHPBonus(item)
}

func recipeHPBonus(item string) int {
	switch item {
	case items.NeuralVisor:
		return 10

	case items.FirewallJacket:
		return 25

	case items.ProxyBoots:
		return 15
	}

	return 0
}