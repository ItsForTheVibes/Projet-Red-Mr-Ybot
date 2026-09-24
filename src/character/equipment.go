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

	default:
		fmt.Println("[-] This item cannot be equipped.")
		return
	}

	// Remove the new equipment from inventory first.
	// This creates space for the old equipment if the inventory was full.
	RemoveItem(c, item, 1)

	// If something is already equipped in this slot,
	// remove its HP bonus and return it to inventory.
	if *slot != "" {
		oldItem := *slot

		c.MaxHP -= equipmentBonus(oldItem)

		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}

		AddItem(c, oldItem)

		fmt.Printf(
			"[>] %s returned to inventory.\n",
			oldItem,
		)
	}

	// Equip the new item.
	*slot = item

	bonus := equipmentBonus(item)

	c.MaxHP += bonus

	fmt.Printf("[+] Equipped: %s\n", item)
	fmt.Printf("[+] Maximum HP +%d\n", bonus)
	fmt.Printf("[+] HP: %d/%d\n", c.CurrentHP, c.MaxHP)
}

func equipmentBonus(item string) int {
	switch item {
	case items.NeuralVisor:
		return 10

	case items.FirewallJacket:
		return 25

	case items.ProxyBoots:
		return 15

	default:
		return 0
	}
}