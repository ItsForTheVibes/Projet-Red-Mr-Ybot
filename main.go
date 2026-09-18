package main

import (
	"Projet-Red/src/character"
	"Projet-Red/src/utils"
)

func main() {
	c := character.CharacterCreation()

	utils.Menu(*c)
}
