package character


type Character struct {
	Name      string   
	Class     string   
	Level     int      
	MaxHP     int      
	CurrentHP int     
	Inventory []string 
}

func InitCharacter (name string,class string,level int,maxhp int,currenthp int,inventory []string) Character{
    return Character{
        Name : name,
        Class : class,
        Level : level,
        MaxHP : maxhp,
        CurrentHP : currenthp,
        Inventory : inventory,
    }
}