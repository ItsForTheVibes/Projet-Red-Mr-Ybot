package main

import (
	"Projet-Red/src/character"
	"Projet-Red/src/utils"
)

func main() {
	utils.Dialogue()

	c := character.CharacterCreation()

	utils.Menu(c)
}
