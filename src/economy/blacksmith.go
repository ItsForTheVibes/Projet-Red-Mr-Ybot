package economy

import (
	"Projet-Red/src/character"
	"Projet-Red/src/items"
	"fmt"
)

func Blacksmith(c *character.Character) {
	for {
		recipes := items.Recipes()

		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║ HARDWARE LAB // FABRICATION TERMINAL        ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println()

		fmt.Printf("Wallet  : %d ETH\n", c.Ethereum)
		fmt.Printf(
			"Storage : %d/%d\n",
			len(c.Inventory),
			c.InventoryCapacity,
		)

		fmt.Println()
		fmt.Println("AVAILABLE BLUEPRINTS")
		fmt.Println()

		for i, recipe := range recipes {
			fmt.Printf(
				"[%d] %s // %d ETH\n",
				i+1,
				recipe.Name,
				recipe.Price,
			)

			fmt.Printf(
				"    %s\n",
				recipe.Description,
			)

			fmt.Println("    Requirements:")

			for _, material := range recipe.Materials {
				owned := character.CountItem(
					c,
					material.Name,
				)

				fmt.Printf(
					"      - %s x%d // owned: %d\n",
					material.Name,
					material.Quantity,
					owned,
				)
			}

			fmt.Println()
		}

		fmt.Println("[0] Leave hardware lab")
		fmt.Println()

		var choice int

		fmt.Print("root@forge:~$ ")
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(recipes) {
			fmt.Println("[-] Invalid blueprint.")
			continue
		}

		craftItem(c, recipes[choice-1])
	}
}

func craftItem(
	c *character.Character,
	recipe items.Recipe,
) {
	if c.Ethereum < recipe.Price {
		fmt.Printf(
			"[-] Need %d ETH to fabricate this equipment.\n",
			recipe.Price,
		)

		return
	}

	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Println("[-] Not enough storage space.")
		return
	}

	for _, material := range recipe.Materials {
		owned := character.CountItem(
			c,
			material.Name,
		)

		if owned < material.Quantity {
			fmt.Printf(
				"[-] Missing %s: %d/%d available.\n",
				material.Name,
				owned,
				material.Quantity,
			)

			return
		}
	}

	for _, material := range recipe.Materials {
		character.RemoveItem(
			c,
			material.Name,
			material.Quantity,
		)
	}

	c.Ethereum -= recipe.Price

	character.AddItem(c, recipe.Name)

	fmt.Println()
	fmt.Println("[+] FABRICATION COMPLETE")
	fmt.Printf("[+] Equipment: %s\n", recipe.Name)
	fmt.Printf("[-] Cost: %d ETH\n", recipe.Price)
	fmt.Printf("[+] Wallet: %d ETH\n", c.Ethereum)
}