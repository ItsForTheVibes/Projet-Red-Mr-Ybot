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
	Ethereum 	  int  
	Skill []string
	Inventory []string 
}


func InitCharacter(name string, class string, maxhp int, skill []string) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     1,                             
		MaxHP:     maxhp,
		CurrentHP: maxhp / 2,
		Ethereum:      100,                    
		Inventory: []string{"Potion", "Potion"},
		Skill:      skill, 
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
		fmt.Println("1. Humain (100 PV)")
		fmt.Println("2. Elfe (80 PV)")
		fmt.Println("3. Nain (120 PV)")
		fmt.Print("Choix (1-3) : ")
		fmt.Scanln(&choice)

		if choice == 1 {
			class = "Humain"
			maxHP = 100
			break
		} else if choice == 2 {
			class = "Elfe"
			maxHP = 80
			break
		} else if choice == 3 {
			class = "Nain"
			maxHP = 120
			break
		}
		fmt.Println("Choix invalide !")
	}

	
	return InitCharacter(name, class, maxHP,[]string{"Coup de poing"})
}


