package character

import (
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Ethereum  int
	Skill     []string
	Inventory []string
	Equipment Equipment
}

func InitCharacter(name string, class string, maxhp int, skill []string) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     1,
		MaxHP:     maxhp,
		CurrentHP: maxhp / 2,
		Ethereum:  100,
		Inventory: []string{"Potion", "Potion"},
		Skill:     skill,
		Equipment: Equipment{
			Head:  "None",
			Torso: "None",
			Feet:  "None",
		},
	}
}

func CharacterCreation() *Character {
	var name string
	var choice int
	var class string
	var maxHP int

	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		fmt.Scanln(&name)

		valid := len(name) > 0
		for _, char := range name {
			if !unicode.IsLetter(char) {
				valid = false
				break
			}
		}

		if valid {
			name = strings.ToLower(name)
			runes := []rune(name)
			runes[0] = unicode.ToUpper(runes[0])
			name = string(runes)
			break
		}
		fmt.Println("Erreur : le nom ne doit contenir que des lettres !")
	}

	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println()

		fmt.Println("1. Anonymous (120 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Freedom of Knowledge / Worldwide")
		fmt.Println()

		fmt.Println("2. LulzSec (100 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Troll / UK-USA")
		fmt.Println()

		fmt.Println("3. Lazarus (110 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Malicious / North Korea")
		fmt.Println()

		fmt.Println("4. White-Hat (90 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Ethical / Other")
		fmt.Println()

		fmt.Println("5. Black-Hat (95 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Malicious / Other")
		fmt.Println()

		fmt.Println("6. Script-Kiddie (60 HP)")
		fmt.Println("   Affinity/Nationality:")
		fmt.Println("   Beginner / Other")
		fmt.Println()

		fmt.Print("Choix (1-6) : ")
		fmt.Scanln(&choice)

		if choice == 1 {
			class = "Anonymous"
			maxHP = 120
			break
		} else if choice == 2 {
			class = "LulzSec"
			maxHP = 100
			break
		} else if choice == 3 {
			class = "Lazarus"
			maxHP = 110
			break
		} else if choice == 4 {
			class = "White-Hat"
			maxHP = 90
			break
		} else if choice == 5 {
			class = "Black-Hat"
			maxHP = 95
			break
		} else if choice == 6 {
			class = "Script-Kiddie"
			maxHP = 60
			break
		}

		fmt.Println("Choix invalide !")
	}

	return InitCharacter(name, class, maxHP, []string{"Coup de poing"})
}
