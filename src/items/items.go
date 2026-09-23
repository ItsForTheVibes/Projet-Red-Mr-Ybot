package items

const (
	// Consumables
	AntiVirusPatch   = "AntiVirus Patch"
	CorruptionScript = "Corruption Script"

	// Skill book
	ExploitManual = "Exploit Manual: Zero-Day Blast"

	// Crafting materials
	EncryptedData  = "Encrypted Data"
	FirewallModule = "Firewall Module"
	SecurityToken  = "Security Token"
	APIKey         = "API Key"

	// Inventory upgrade
	StorageExpansion = "Storage Expansion Module"

	// Equipment
	NeuralVisor    = "Neural Visor"
	FirewallJacket = "Firewall Jacket"
	ProxyBoots     = "Proxy Boots"

	// Skills
	PacketPunch  = "Packet Punch"
	ZeroDayBlast = "Zero-Day Blast"
)

type ShopItem struct {
	Name        string
	Description string
	Price       int
}

type Material struct {
	Name     string
	Quantity int
}

type Recipe struct {
	Name        string
	Description string
	Price       int
	Materials   []Material
}

func MerchantStock() []ShopItem {
	return []ShopItem{
		{
			Name:        AntiVirusPatch,
			Description: "Recovery payload // restores 50 HP",
			Price:       3,
		},
		{
			Name:        CorruptionScript,
			Description: "Offensive payload // 10 damage/sec for 3 seconds",
			Price:       6,
		},
		{
			Name:        ExploitManual,
			Description: "Skill manual // teaches Zero-Day Blast",
			Price:       25,
		},
		{
			Name:        EncryptedData,
			Description: "Crafting material // used for armor and boots",
			Price:       4,
		},
		{
			Name:        FirewallModule,
			Description: "Crafting material // used for chest armor",
			Price:       7,
		},
		{
			Name:        SecurityToken,
			Description: "Crafting material // used for visor and boots",
			Price:       3,
		},
		{
			Name:        APIKey,
			Description: "Crafting material // used for visor",
			Price:       1,
		},
		{
			Name:        StorageExpansion,
			Description: "Inventory upgrade // +10 slots, maximum 3 upgrades",
			Price:       30,
		},
	}
}

func Recipes() []Recipe {
	return []Recipe{
		{
			Name:        NeuralVisor,
			Description: "HEAD // +10 maximum HP",
			Price:       5,
			Materials: []Material{
				{Name: APIKey, Quantity: 1},
				{Name: SecurityToken, Quantity: 1},
			},
		},
		{
			Name:        FirewallJacket,
			Description: "CHEST // +25 maximum HP",
			Price:       5,
			Materials: []Material{
				{Name: EncryptedData, Quantity: 2},
				{Name: FirewallModule, Quantity: 1},
			},
		},
		{
			Name:        ProxyBoots,
			Description: "FEET // +15 maximum HP",
			Price:       5,
			Materials: []Material{
				{Name: EncryptedData, Quantity: 1},
				{Name: SecurityToken, Quantity: 1},
			},
		},
	}
}

func FindRecipe(name string) (Recipe, bool) {
	for _, recipe := range Recipes() {
		if recipe.Name == name {
			return recipe, true
		}
	}

	return Recipe{}, false
}