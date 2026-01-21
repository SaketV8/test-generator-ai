package utils

import (
	"fmt"
	"strings"
)

func PrintBoxedText(title string) {
	padding := 2
	content := strings.Repeat(" ", padding) + title + strings.Repeat(" ", padding)
	border := "+" + strings.Repeat("-", len(content)) + "+"

	fmt.Println()
	fmt.Println(border)
	fmt.Printf("|%s|\n", content)
	fmt.Println(border)
	fmt.Println()
}
