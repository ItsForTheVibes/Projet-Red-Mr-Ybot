package character

// Structure répertoriant toutes les caractéristiques du personnage (Tâche 1)
type Character struct {
	Name      string   // Nom du personnage (Chaîne de caractères)
	Class     string   // Classe du personnage (ex: Elfe, Humain, Nain)
	Level     int      // Niveau actuel du personnage (Nombre entier)
	MaxHP     int      // Points de vie maximum (Nombre entier)
	CurrentHP int      // Points de vie actuels (Nombre entier)
	Inventory []string // Inventaire sous forme de liste de chaînes de caractères (slice)
}