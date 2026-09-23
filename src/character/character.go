package character

import (
	"Projet-Red/src/items"
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Name              string
	Class             string
	Level             int
	MaxHP             int
	CurrentHP         int
	Ethereum          int
	Skill             []string
	Inventory         []string
	InventoryCapacity int
	InventoryUpgrades int
	Equipment         Equipment
}

func InitCharacter(name string, class string, maxHP int) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     1,
		MaxHP:     maxHP,
		CurrentHP: maxHP / 2,
		Ethereum:  100,
		Skill:     []string{items.PacketPunch},
		Inventory: []string{
			items.AntiVirusPatch,
			items.AntiVirusPatch,
			items.AntiVirusPatch,
		},
		InventoryCapacity: 10,
		InventoryUpgrades: 0,
		Equipment:         Equipment{},
	}
}

func CharacterCreation() *Character {
	var name string
	var choice int

	for {
		fmt.Print("Enter operator alias: ")
		fmt.Scanln(&name)

		valid := len(name) > 0

		for _, char := range name {
			if !unicode.IsLetter(char) {
				valid = false
			}
		}

		if valid {
			name = strings.ToLower(name)

			runes := []rune(name)
			runes[0] = unicode.ToUpper(runes[0])

			name = string(runes)
			break
		}

		fmt.Println("[-] Alias must contain letters only.")
	}

	for {
		fmt.Println()
		fmt.Println("=== SELECT AFFILIATION ===")
		fmt.Println()
		fmt.Println("[1] Anonymous  // 100 HP")
		fmt.Println("[2] LulzSec    // 80 HP")
		fmt.Println("[3] Lazarus    // 120 HP")
		fmt.Println()

		fmt.Print("Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			return InitCharacter(name, "Anonymous", 100)

		case 2:
			return InitCharacter(name, "LulzSec", 80)

		case 3:
			return InitCharacter(name, "Lazarus", 120)

		default:
			fmt.Println("[-] Invalid choice.")
		}
	}
}

func IsDead(c *Character) bool {
	if c.CurrentHP > 0 {
		return false
	}

	fmt.Println()
	fmt.Println("[-] OPERATOR CONNECTION LOST")

	c.CurrentHP = c.MaxHP / 2

	fmt.Printf(
		"[+] Emergency restore complete: %d/%d HP\n",
		c.CurrentHP,
		c.MaxHP,
	)

	return true
}
