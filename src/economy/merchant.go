package economy

import (
	"Projet-Red/src/character"
	"Projet-Red/src/items"
	"fmt"
)

func Merchant(c *character.Character) {
	for {
		stock := items.MerchantStock()

		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║ DARKNET // PAYLOAD MARKET                   ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println()

		fmt.Printf("Wallet  : %d ETH\n", c.Ethereum)
		fmt.Printf(
			"Storage : %d/%d\n",
			len(c.Inventory),
			c.InventoryCapacity,
		)

		fmt.Println()
		fmt.Println("AVAILABLE PACKAGES")
		fmt.Println()

		for i, item := range stock {
			fmt.Printf(
				"[%d] %s // %d ETH\n",
				i+1,
				item.Name,
				item.Price,
			)

			fmt.Printf(
				"    %s\n\n",
				item.Description,
			)
		}

		fmt.Println("[0] Disconnect from market")
		fmt.Println()

		var choice int

		fmt.Print("root@darknet:~$ ")
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(stock) {
			fmt.Println("[-] Invalid selection.")
			continue
		}

		buyItem(c, stock[choice-1])
	}
}

func buyItem(
	c *character.Character,
	item items.ShopItem,
) {
	if c.Ethereum < item.Price {
		fmt.Printf(
			"[-] Insufficient Ethereum. Need %d ETH.\n",
			item.Price,
		)

		return
	}

	if len(c.Inventory) >= c.InventoryCapacity {
		fmt.Println("[-] Storage full.")
		fmt.Println("[!] Buy/use a Storage Expansion Module.")

		return
	}

	c.Ethereum -= item.Price

	character.AddItem(c, item.Name)

	fmt.Printf(
		"[+] Downloaded: %s\n",
		item.Name,
	)

	fmt.Printf(
		"[-] Transaction: %d ETH\n",
		item.Price,
	)

	fmt.Printf(
		"[+] Wallet: %d ETH\n",
		c.Ethereum,
	)
}