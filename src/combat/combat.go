package combat

import (
	"Projet-Red/src/character"
	"Projet-Red/src/items"
	"fmt"
	"time"
)

func TrainingFight(c *character.Character) {
	enemy := InitTrainingBot()
	turn := 1

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║ TRAINING NETWORK // COMBAT SIMULATION       ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Printf(
		"[!] Hostile process detected: %s\n",
		enemy.Name,
	)

	for c.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Println()
		fmt.Printf(
			"========== TURN %d ==========\n",
			turn,
		)

		displayCombatStatus(c, enemy)

		if !CharacterTurn(c, enemy) {
			continue
		}

		if enemy.CurrentHP <= 0 {
			fmt.Println()
			fmt.Printf(
				"[+] %s terminated.\n",
				enemy.Name,
			)

			fmt.Println("[+] Training simulation complete.")
			return
		}

		GoblinPattern(c, enemy, turn)

		if character.IsDead(c) {
			return
		}

		turn++
	}
}

func CharacterTurn(
	c *character.Character,
	enemy *Monster,
) bool {
	fmt.Println()
	fmt.Println("PLAYER ACTION")
	fmt.Println("[1] Basic Attack")
	fmt.Println("[2] Inventory")
	fmt.Println("[3] Skills")
	fmt.Println()

	var choice int

	fmt.Print("root@combat:~$ ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		damage := 5

		enemy.CurrentHP -= damage

		if enemy.CurrentHP < 0 {
			enemy.CurrentHP = 0
		}

		fmt.Printf(
			"[>] Basic Attack deals %d damage to %s.\n",
			damage,
			enemy.Name,
		)

		fmt.Printf(
			"[>] %s HP: %d/%d\n",
			enemy.Name,
			enemy.CurrentHP,
			enemy.MaxHP,
		)

		return true

	case 2:
		return combatInventory(c, enemy)

	case 3:
		return useSkill(c, enemy)

	default:
		fmt.Println("[-] Invalid action.")
		return false
	}
}

func GoblinPattern(
	c *character.Character,
	enemy *Monster,
	turn int,
) {
	damage := enemy.Attack

	if turn%3 == 0 {
		damage = enemy.Attack * 2

		fmt.Println()
		fmt.Println("[!] CRITICAL ATTACK PATTERN DETECTED")
	}

	c.CurrentHP -= damage

	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	fmt.Printf(
		"[<] %s attacks %s for %d damage.\n",
		enemy.Name,
		c.Name,
		damage,
	)

	fmt.Printf(
		"[<] %s HP: %d/%d\n",
		c.Name,
		c.CurrentHP,
		c.MaxHP,
	)
}

func combatInventory(
	c *character.Character,
	enemy *Monster,
) bool {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║ COMBAT // STORAGE ACCESS                    ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Println()

	if len(c.Inventory) == 0 {
		fmt.Println("[!] Inventory empty.")
		return false
	}

	for i, item := range c.Inventory {
		fmt.Printf(
			"[%d] %s\n",
			i+1,
			item,
		)
	}

	fmt.Println()
	fmt.Println("[0] Return")
	fmt.Println()

	var choice int

	fmt.Print("root@combat-storage:~$ ")
	fmt.Scanln(&choice)

	if choice == 0 {
		return false
	}

	if choice < 1 || choice > len(c.Inventory) {
		fmt.Println("[-] Invalid item.")
		return false
	}

	item := c.Inventory[choice-1]

	switch item {
	case items.AntiVirusPatch:
		character.TakePot(c)
		return true

	case items.CorruptionScript:
		character.RemoveItem(
			c,
			items.CorruptionScript,
			1,
		)

		PoisonPot(enemy)

		return true

	default:
		fmt.Printf(
			"[!] %s cannot be used during combat.\n",
			item,
		)

		return false
	}
}

func PoisonPot(enemy *Monster) {
	fmt.Printf(
		"[>] Corruption Script injected into %s.\n",
		enemy.Name,
	)

	for i := 1; i <= 3; i++ {
		enemy.CurrentHP -= 10

		if enemy.CurrentHP < 0 {
			enemy.CurrentHP = 0
		}

		fmt.Printf(
			"[>] Corruption %d/3 // %s HP: %d/%d\n",
			i,
			enemy.Name,
			enemy.CurrentHP,
			enemy.MaxHP,
		)

		if enemy.CurrentHP <= 0 {
			return
		}

		time.Sleep(time.Second)
	}
}

func useSkill(
	c *character.Character,
	enemy *Monster,
) bool {
	fmt.Println()
	fmt.Println("=== INSTALLED EXPLOITS ===")

	for i, skill := range c.Skill {
		fmt.Printf(
			"[%d] %s\n",
			i+1,
			skill,
		)
	}

	fmt.Println("[0] Return")

	var choice int

	fmt.Print("root@exploit:~$ ")
	fmt.Scanln(&choice)

	if choice == 0 {
		return false
	}

	if choice < 1 || choice > len(c.Skill) {
		fmt.Println("[-] Invalid exploit.")
		return false
	}

	skill := c.Skill[choice-1]

	damage := 0

	switch skill {
	case items.PacketPunch:
		damage = 8

	case items.ZeroDayBlast:
		damage = 18

	default:
		fmt.Println("[-] Unknown exploit.")
		return false
	}

	enemy.CurrentHP -= damage

	if enemy.CurrentHP < 0 {
		enemy.CurrentHP = 0
	}

	fmt.Printf(
		"[>] %s executed // %d damage.\n",
		skill,
		damage,
	)

	fmt.Printf(
		"[>] %s HP: %d/%d\n",
		enemy.Name,
		enemy.CurrentHP,
		enemy.MaxHP,
	)

	return true
}

func displayCombatStatus(
	c *character.Character,
	enemy *Monster,
) {
	fmt.Println()
	fmt.Printf(
		"OPERATOR // %s // HP %d/%d\n",
		c.Name,
		c.CurrentHP,
		c.MaxHP,
	)

	fmt.Printf(
		"TARGET   // %s // HP %d/%d\n",
		enemy.Name,
		enemy.CurrentHP,
		enemy.MaxHP,
	)
}
