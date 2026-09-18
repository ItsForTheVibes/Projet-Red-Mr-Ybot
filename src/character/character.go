package character


type Character struct {
	Name      string   
	Class     string   
	Level     int      
	MaxHP     int      
	CurrentHP int     
	Inventory []string 
}


func InitCharacter(name string, class string, maxhp int) *Character {
	return &Character{
		Name:      name,
		Class:     class,
		Level:     1,                             
		MaxHP:     maxhp,
		CurrentHP: maxhp / 2,                    
		Inventory: []string{"Potion", "Potion"}, 
	}
}
