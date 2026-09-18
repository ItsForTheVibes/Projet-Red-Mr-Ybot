package combat

import ("fmt"
		"time"
)

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


func poisonPot(player *Player) {
    for i := 0; i < 3; i++ {
        player.CurrentHP -= 10

        if player.CurrentHP < 0 {
            player.CurrentHP = 0
        }

        fmt.Printf("Points de vie : %d/%d\n", player.CurrentHP, player.MaxHP)

        time.Sleep(1 * time.Second)
    }
}