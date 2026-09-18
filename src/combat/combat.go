package combat

import ("fmt"
		"time"
		"Projet-Red/src/character"
)

func takePot(character *character.Character) { // Variable, package et structure

	for i, item := range character.Inventory {
		if item == "potion" {
			character.Inventory = append(character.Inventory[:i], character.Inventory[i+1:]...)

			character.CurrentHP += 50

			if character.CurrentHP > character.MaxHP {
				character.CurrentHP = character.MaxHP
			}

			fmt.Printf("Vous avez utilisez la potion !\n Points de vie : %d/%d\n", character.CurrentHP, character.MaxHP)

			return
		}
	}
	println("Vous n'avez pas de potion.")
}


func poisonPot(character *character.Character) {
    for i := 0; i < 3; i++ {
        character.CurrentHP -= 10

        if character.CurrentHP < 0 {
            character.CurrentHP = 0
        }

        fmt.Printf("Points de vie : %d/%d\n", character.CurrentHP, character.MaxHP)

        time.Sleep(1 * time.Second)
    }
}

func spellBook(character *character.Character) {
	for _, skill := range character.Skill {
		if skill == "Boule de Feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de Feu.")
			return
		}
	}

	character.Skill = append(character.Skill, "Boule de Feu")
	fmt.Println("Vous avez appris le sort Boule de Feu !")
}