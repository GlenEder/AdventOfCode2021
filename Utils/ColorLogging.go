package Utils

import "fmt"

const (
	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorGreen = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan = "\033[36m"
	colorWhite = "\033[37m"

	ColorReset = 0
	Red = 1
	Green = 2
	Yellow = 3
	Blue = 4
	Purple = 5
	Cyan = 6
	White = 7
)

//Prints the string provided in color provided
//@param c -- color to print in
//@param s -- string to print
func PrintWithColor(c int, s string) {
	SetOutputColor(c)
	fmt.Printf("%s", s)
	ResetOutputColor()
}

//Prints provided string in red color
//@param s -- error to print
func PrintError(s string) {
	SetOutputColor(Red)
	fmt.Printf("%s", s)
	ResetOutputColor()
}

//Sets output color in terminal
//@param c -- color int to use (see const defined at top of utils.go)
func SetOutputColor(c int)  {
	switch c {
	case Red:
		fmt.Print(colorRed)
		break
	case Green:
		fmt.Print(colorGreen)
		break
	case Yellow:
		fmt.Print(colorYellow)
		break
	case Blue:
		fmt.Print(colorBlue)
		break
	case Purple:
		fmt.Print(colorPurple)
		break
	case Cyan:
		fmt.Print(colorCyan)
		break
	case White:
		fmt.Print(colorWhite)
		break
	}
}

//Returns color code for setting terminal output color
//@param c -- color int to use (see const defined at top of utils.go)
//@return color code for terminal output
func GetColorCode(c int) string {
	switch c {
	case ColorReset:
		return colorReset
	case Red:
		return colorRed
	case Green:
		return colorGreen
	case Yellow:
		return colorYellow
	case Blue:
		return colorBlue
	case Purple:
		return colorPurple
	case Cyan:
		return colorCyan
	case White:
		return colorWhite
	}

	return ""
}

func ResetOutputColor() {
	fmt.Print(colorReset)
}