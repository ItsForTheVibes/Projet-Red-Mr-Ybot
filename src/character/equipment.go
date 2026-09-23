package character

import "fmt"

// Structure Equipment répertoriant les 3 emplacements d'équipement
type Equipment struct {
	Head  string // Équipement de tête
	Chest string // Équipement pour le torse
	Feet  string // Équipement pour les pieds
}

// EquipItem tente d'équiper un objet présent dans l'inventaire du personnage
func (c *Character) EquipItem(itemName string) {
	// Vérifier si l'objet est bien dans l'inventaire
	itemIndex := -1
	for i, item := range c.Inventory {
		if item == itemName {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		fmt.Printf("Vous ne possédez pas '%s' dans votre inventaire.\n", itemName)
		return
	}

	var slot *string
	var hpBonus int

	// Déterminer le slot concerné et le bonus de PV max associés à l'équipement
	switch itemName {
	case "Chapeau de l'aventurier":
		slot = &c.Equipment.Head
		hpBonus = 10
	case "Tunique de l'aventurier":
		slot = &c.Equipment.Chest
		hpBonus = 25
	case "Bottes de l'aventurier":
		slot = &c.Equipment.Feet
		hpBonus = 15
	default:
		fmt.Printf("'%s' n'est pas un équipement équipable.\n", itemName)
		return
	}

	// Retirer l'objet à équiper de l'inventaire
	c.Inventory = append(c.Inventory[:itemIndex], c.Inventory[itemIndex+1:]...)

	// Si un équipement est déjà équipé sur ce slot, on le déséquipe
	if *slot != "" {
		oldItem := *slot
		oldHpBonus := getEquipmentHPBonus(oldItem)

		// Retirer l'ancien bonus de PV max
		c.MaxHP -= oldHpBonus
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}

		// Remettre l'ancien équipement dans l'inventaire
		c.Inventory = append(c.Inventory, oldItem)
		fmt.Printf("🔄 Vous déséquipez '%s' et le remettez dans l'inventaire.\n", oldItem)
	}

	// Équiper le nouvel objet et appliquer le bonus de PV max
	*slot = itemName
	c.MaxHP += hpBonus
	fmt.Printf("Vous avez équipé '%s' ! (+%d PV max)\n", itemName, hpBonus)
	fmt.Printf("PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
}

// Fonction utilitaire pour connaître le bonus de PV d'un équipement
func getEquipmentHPBonus(itemName string) int {
	switch itemName {
	case "Chapeau de l'aventurier":
		return 10
	case "Tunique de l'aventurier":
		return 25
	case "Bottes de l'aventurier":
		return 15
	default:
		return 0
	}
}