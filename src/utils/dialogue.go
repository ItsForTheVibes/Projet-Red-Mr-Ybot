package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dialogue() {
	fmt.Println(`
                  ██████╗ ██████╗  ██████╗      ██╗███████╗ ██████╗████████╗
                  ██╔══██╗██╔══██╗██╔═══██╗     ██║██╔════╝██╔════╝╚══██╔══╝
                  ██████╔╝██████╔╝██║   ██║     ██║█████╗  ██║        ██║
                  ██╔═══╝ ██╔══██╗██║   ██║██   ██║██╔══╝  ██║        ██║
                  ██║     ██║  ██║╚██████╔╝╚█████╔╝███████╗╚██████╗   ██║
                  ╚═╝     ╚═╝  ╚═╝ ╚═════╝  ╚════╝ ╚══════╝ ╚═════╝   ╚═╝

                                ██████╗ ███████╗██████╗
                                ██╔══██╗██╔════╝██╔══██╗
                                ██████╔╝█████╗  ██║  ██║
                                ██╔══██╗██╔══╝  ██║  ██║
                                ██║  ██║███████╗██████╔╝
                                ╚═╝  ╚═╝╚══════╝╚═════╝

                  ┌──────────────────────────────────────────────────────┐
                  │              R E S T R I C T E D   N E T             │
                  │                                                      │
                  │        > Connexion au réseau clandestin... OK        │
                  │        > Chiffrement de la connexion......... OK     │
                  │        > Identification.................... INCONNUE │
                  │                                                      │
                  │                 [ ACCÈS AUTORISÉ ]                   │
                  └──────────────────────────────────────────────────────┘
	`)
	reader := bufio.NewReader(os.Stdin)

	dialogue := []string{
		"[SYSTÈME] Bienvenue, recrue. Le réseau est bien plus vaste que tu ne l'imagines... et bien plus dangereux.",
		"[INCONNU] Ici, ton vrai nom ne vaut rien. Il te faut une nouvelle identité.",
		"[INCONNU] Choisis ton pseudo, rejoins un groupe de hackers et forge-toi une réputation.",
	}

	fmt.Println("=== PROJECT RED ===")
	fmt.Println("Appuie sur ENTRÉE pour continuer ou écris 'skip' pour passer l'introduction.")
	fmt.Println()

	for _, line := range dialogue {
		fmt.Println(line)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "skip") {
			break
		}
	}

	fmt.Println()
	fmt.Println("[SYSTÈME] Il est temps de créer ton identité.")
	fmt.Println()
}
