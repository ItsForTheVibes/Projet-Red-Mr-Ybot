package utils

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Header(title string) {
	fmt.Println(Cyan + "╔══════════════════════════════════════════════╗" + Reset)
	fmt.Printf(Cyan+"║"+Reset+" %-44s "+Cyan+"║\n"+Reset, title)
	fmt.Println(Cyan + "╚══════════════════════════════════════════════╝" + Reset)
}

func ClassHeader(title string, class string) {
	color := ClassColor(class)

	fmt.Println(color + "╔══════════════════════════════════════════════╗" + Reset)
	fmt.Printf(color+"║"+Reset+" %-44s "+color+"║\n"+Reset, title)
	fmt.Println(color + "╚══════════════════════════════════════════════╝" + Reset)
}

func Success(message string) {
	fmt.Println(Green + "[+] " + Reset + message)
}

func Warning(message string) {
	fmt.Println(Yellow + "[!] " + Reset + message)
}

func Error(message string) {
	fmt.Println(Red + "[-] " + Reset + message)
}

func Prompt() {
	fmt.Print(Green + "root@project-red" + Reset + ":" + Blue + "~$ " + Reset)
}

func HealthBar(current int, max int) string {
	width := 20

	if max <= 0 {
		return "[????????????????????]"
	}

	filled := current * width / max

	if filled < 0 {
		filled = 0
	}

	if filled > width {
		filled = width
	}

	return "[" +
		strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled) +
		"]"
}

func ClassColor(class string) string {
	switch class {
	case "Anonymous":
		return "\033[38;2;0;255;100m"

	case "LulzSec":
		return "\033[38;2;200;0;255m"

	case "Lazarus":
		return "\033[38;2;255;30;30m"

	case "White-Hat":
		return "\033[38;2;0;220;255m"

	case "Black-Hat":
		return "\033[38;2;180;0;0m"

	case "Script-Kiddie":
		return "\033[38;2;255;200;0m"

	default:
		return White
	}
}
