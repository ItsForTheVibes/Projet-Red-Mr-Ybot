package main

import (
	//"Projet-Red/economy"
	"Projet-Red/src/character"
	"Projet-Red/src/utils"
	//"fmt"
)

func main() {
	c := character.CharacterCreation()
	utils.Menu(*c)

}
