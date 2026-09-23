package character

import (
	"Projet-Red/src/items"
	"fmt"
)

func AccessInventory(c *Character) {
	for {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║ STORAGE // OPERATOR INVENTORY               ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println()

		fmt.Printf(
			"Capacity: %d/%d\n",
			len(c.Inventory),
			c.InventoryCapacity,
		)

		fmt.Println()

		if len(c.Inventory) == 0 {
			fmt.Println("[ EMPTY ]")
			return
		}

		for i, item := range c.Inventory {
			fmt.Printf("[%02d] %s\n", i+1, item)
		}

		fmt.Println()
		fmt.Println("[0] Return")
		fmt.Println()

		var choice int

		fmt.Print("root@storage:~$ ")
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("[-] Invalid selection.")
			continue
		}

		item := c.Inventory[choice-1]

		UseItem(c, item)
	}
}

func UseItem(c *Character, item string) bool {
	switch item {
	case items.AntiVirusPatch:
		TakePot(c)
		return true

	case items.ExploitManual:
		return SpellBook(c)

	case items.StorageExpansion:
		return UpgradeInventorySlot(c)

	case items.NeuralVisor,
		items.FirewallJacket,
		items.ProxyBoots:

		EquipItem(c, item)
		return true

	default:
		fmt.Println("[!] This item cannot be used here.")
		return false
	}
}

func TakePot(c *Character) {
	if !RemoveItem(c, items.AntiVirusPatch, 1) {
		fmt.Println("[-] No AntiVirus Patch available.")
		return
	}

	c.CurrentHP += 50

	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Println("[+] AntiVirus Patch executed.")
	fmt.Printf("[+] HP: %d/%d\n", c.CurrentHP, c.MaxHP)
}

func SpellBook(c *Character) bool {
	for _, skill := range c.Skill {
		if skill == items.ZeroDayBlast {
			fmt.Println("[!] Zero-Day Blast is already installed.")
			return false
		}
	}

	c.Skill = append(c.Skill, items.ZeroDayBlast)

	RemoveItem(c, items.ExploitManual, 1)

	fmt.Println("[+] Exploit installed.")
	fmt.Println("[+] New skill unlocked: Zero-Day Blast")

	return true
}

func UpgradeInventorySlot(c *Character) bool {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("[-] Storage is already at maximum capacity.")
		return false
	}

	if !RemoveItem(c, items.StorageExpansion, 1) {
		fmt.Println("[-] Storage Expansion Module not found.")
		return false
	}

	c.InventoryCapacity += 10
	c.InventoryUpgrades++

	fmt.Printf(
		"[+] Storage expanded to %d slots.\n",
		c.InventoryCapacity,
	)

	fmt.Printf(
		"[+] Upgrades used: %d/3\n",
		c.InventoryUpgrades,
	)

	return true
}

func AddItem(c *Character, item string) bool {
	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Println("[-] Inventory full.")
		return false
	}

	c.Inventory = append(c.Inventory, item)

	return true
}

func RemoveItem(c *Character, item string, quantity int) bool {
	if CountItem(c, item) < quantity {
		return false
	}

	for removed := 0; removed < quantity; removed++ {
		for i, inventoryItem := range c.Inventory {
			if inventoryItem == item {
				c.Inventory = append(
					c.Inventory[:i],
					c.Inventory[i+1:]...,
				)

				break
			}
		}
	}

	return true
}

func CountItem(c *Character, item string) int {
	count := 0

	for _, inventoryItem := range c.Inventory {
		if inventoryItem == item {
			count++
		}
	}

	return count
}
