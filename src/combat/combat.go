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


func poisonPot(c *character.Character) {
    for i := 0; i < 3; i++ {
        c.CurrentHP -= 10

        if c.CurrentHP < 0 {
            c.CurrentHP = 0
        }

        fmt.Printf("Points de vie : %d/%d\n", c.CurrentHP, c.MaxHP)

		character.IsDead(c)
        time.Sleep(1 * time.Second)
    }
}
