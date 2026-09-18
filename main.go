package main

import (
	"Projet-Red/src/character"
	"Projet-Red/src/economy"
	"fmt"
)

func main() {
	c := character.Character{}

	economy.Merchant(&c)

	fmt.Println("Inventory:", c.Inventory)

	fmt.Println(utils.Red + "RED TEST" + utils.Reset)
	fmt.Println(utils.Green + "GREEN TEST" + utils.Reset)
	fmt.Println(utils.Cyan + "CYAN TEST" + utils.Reset)
}
