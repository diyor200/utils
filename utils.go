package utils

import (
	"fmt"

	colorize "github.com/fatih/color"
)

const (
	ColorRed = iota + 1
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
)

func PrintColor(text string, color int) {
	switch color {
	case ColorRed:
		colorize.Red(text)
	case ColorGreen:
		colorize.Green(text)
	case ColorYellow:
		colorize.Yellow(text)
	case ColorBlue:
		colorize.Blue(text)
	case ColorMagenta:
		colorize.Magenta(text)
	default:
		fmt.Println(text)
	}
}

func ContainsInt(a []int, x int) {
	for _, n := range a {
		if x == n {
			colorize.Green("Exists")
		}
	}

	colorize.Red("Does not Exists")
}
