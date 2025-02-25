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
	default:
		fmt.Println(text)
	}
}
