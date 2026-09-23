package combat

import (
	"Projet-Red/src/character"
	"fmt"
	"time"
)

func StartCombat(c *character.Character) {
	enemy := NewMonster("Malware Bot", 60, 10)

	fmt.Println()
	fmt.Println("===== COMBAT =====")
	fmt.Printf("Un %s apparaît !\n", enemy.Name)

	for c.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Println()

		fmt.Printf(
			"%s : %d/%d HP\n",
			c.Name,
			c.CurrentHP,
			c.MaxHP,
		)

		fmt.Printf(
			"%s : %d/%d HP\n",
			enemy.Name,
			enemy.CurrentHP,
			enemy.MaxHP,
		)

		fmt.Println()
		fmt.Println("1 - Coup de poing")
		fmt.Println("2 - Potion de vie")
		fmt.Println("3 - Payload")
		fmt.Println("4 - Boule de Feu")
		fmt.Println("5 - Fuir")

		var choice int

		fmt.Print("Choix : ")
		fmt.Scanln(&choice)

		actionTaken := false

		switch choice {
		case 1:
			damage := 10

			enemy.CurrentHP -= damage

			if enemy.CurrentHP < 0 {
				enemy.CurrentHP = 0
			}

			fmt.Printf(
				"Vous infligez %d dégâts à %s.\n",
				damage,
				enemy.Name,
			)

			actionTaken = true

		case 2:
			actionTaken = takePotion(c)

		case 3:
			actionTaken = usePayload(c, enemy)

		case 4:
			actionTaken = useFireball(c, enemy)

		case 5:
			fmt.Println("Vous fuyez le combat.")
			return

		default:
			fmt.Println("Choix invalide.")
		}

		if !actionTaken {
			continue
		}

		if enemy.CurrentHP <= 0 {
			fmt.Printf(
				"\n%s a été vaincu !\n",
				enemy.Name,
			)

			reward := 10

			c.Ethereum += reward

			fmt.Printf(
				"Vous gagnez %d Ethereum.\n",
				reward,
			)

			return
		}

		enemyAttack(c, enemy)

		if character.IsDead(c) {
			return
		}
	}
}

func enemyAttack(c *character.Character, enemy *Monster) {
	c.CurrentHP -= enemy.Attack

	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	fmt.Printf(
		"%s vous attaque et inflige %d dégâts.\n",
		enemy.Name,
		enemy.Attack,
	)
}

func takePotion(c *character.Character) bool {
	for i, item := range c.Inventory {
		if item == "Potion de vie" {
			c.Inventory = append(
				c.Inventory[:i],
				c.Inventory[i+1:]...,
			)

			c.CurrentHP += 50

			if c.CurrentHP > c.MaxHP {
				c.CurrentHP = c.MaxHP
			}

			fmt.Printf(
				"Potion utilisée ! HP : %d/%d\n",
				c.CurrentHP,
				c.MaxHP,
			)

			return true
		}
	}

	fmt.Println("Vous n'avez pas de Potion de vie.")

	return false
}

func usePayload(c *character.Character, enemy *Monster) bool {
	fmt.Println()
	fmt.Println("===== PAYLOADS =====")
	fmt.Println("1 - AntiVirus")
	fmt.Println("2 - Corruption Script")
	fmt.Println("3 - Exploit Script")
	fmt.Println("4 - Retour")

	var choice int

	fmt.Print("Choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		if !removeItem(c, "AntiVirus") {
			fmt.Println("Vous ne possédez pas AntiVirus.")
			return false
		}

		c.CurrentHP += 10

		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}

		fmt.Println("AntiVirus utilisé : +10 HP.")

		return true

	case 2:
		if !removeItem(c, "Corruption Script") {
			fmt.Println("Vous ne possédez pas Corruption Script.")
			return false
		}

		fmt.Println("Corruption Script lancé !")

		poisonPot(enemy)

		return true

	case 3:
		if !removeItem(c, "Exploit Script") {
			fmt.Println("Vous ne possédez pas Exploit Script.")
			return false
		}

		enemy.CurrentHP -= 25

		if enemy.CurrentHP < 0 {
			enemy.CurrentHP = 0
		}

		fmt.Println("Exploit Script inflige 25 dégâts.")

		return true

	case 4:
		return false

	default:
		fmt.Println("Choix invalide.")
		return false
	}
}

func useFireball(c *character.Character, enemy *Monster) bool {
	for _, skill := range c.Skill {
		if skill == "Boule de Feu" {
			damage := 20

			enemy.CurrentHP -= damage

			if enemy.CurrentHP < 0 {
				enemy.CurrentHP = 0
			}

			fmt.Printf(
				"Boule de Feu inflige %d dégâts !\n",
				damage,
			)

			return true
		}
	}

	fmt.Println("Vous ne connaissez pas Boule de Feu.")

	return false
}

func poisonPot(enemy *Monster) {
	for i := 0; i < 3; i++ {
		enemy.CurrentHP -= 10

		if enemy.CurrentHP < 0 {
			enemy.CurrentHP = 0
		}

		fmt.Printf(
			"%s prend 10 dégâts de poison. HP : %d/%d\n",
			enemy.Name,
			enemy.CurrentHP,
			enemy.MaxHP,
		)

		if enemy.CurrentHP <= 0 {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func removeItem(c *character.Character, itemName string) bool {
	for i, item := range c.Inventory {
		if item == itemName {
			c.Inventory = append(
				c.Inventory[:i],
				c.Inventory[i+1:]...,
			)

			return true
		}
	}

	return false
}