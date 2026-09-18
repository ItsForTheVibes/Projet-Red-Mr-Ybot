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
}
