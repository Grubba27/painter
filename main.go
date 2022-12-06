package paint

import "runtime"

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func colorize(text string, color string) string {
	if runtime.GOOS == "windows" {
		return text
	}
	return color + text + reset
}

func InRed(text string) string {
	return colorize(text, red)
}
func InGreen(text string) string {
	return colorize(text, green)
}
func InYellow(text string) string {
	return colorize(text, yellow)
}
func InBlue(text string) string {
	return colorize(text, blue)
}
func InPurple(text string) string {
	return colorize(text, purple)
}
func InCyan(text string) string {
	return colorize(text, cyan)
}
func InGray(text string) string {
	return colorize(text, gray)
}
func InWhite(text string) string {
	return colorize(text, white)
}
