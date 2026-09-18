package combat

import "fmt"

func takePot(player *Player) { // Vérifier qu'il y a une potion dans l'inventaire

	for i, item := range player.Inventory {
		if item == "potion" {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)

			player.CurrentHP += 50

			if player.CurrentHP > player.MaxHP {
				player.CurrentHP = player.MaxHP
			}

			fmt.Printf("Vous avez utilisez la potion !\n Points de vie : %d/%d\n", player.CurrentHP, player.MaxHP)

			return
		}
	}
	println("Vous n'avez pas de potion.")
}
