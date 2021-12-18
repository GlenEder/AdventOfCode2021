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

//Prints the formatted string with color
//@param c -- color to print in
//@param s -- formatted string to print
//@param a -- args for string
func PrintWithColorf(c int, s string, a ...interface{}) {
	SetOutputColor(c)
	fmt.Printf(s, a...)
	ResetOutputColor()
}


//Prints with desired tabs placed in front of string in color
//@param c -- color to print in
//@param depth -- number of tabs to add
//@param s -- string formatting
//@param a -- args for string
func PrintRecursionLogColorf(c int, depth int, s string, a ...interface{}) {
	for i := 0; i < depth; i++ {
		fmt.Printf("\t")
	}
	PrintWithColorf(c, s, a...)
}

//Prints provided string in red color
//@param s -- error to print
func PrintError(s string) {
	SetOutputColor(Red)
	fmt.Printf("%s", s)
	ResetOutputColor()
}


func PrintGridWithHighlights(grid [][]int, highlights []Point) {
	for y, line := range grid {
		for x, num := range line {
			p := Point{X: x, Y: y}
			if p.InPointSlice(highlights) {
				PrintWithColorf(Green, "%d", num)
			} else {
				fmt.Printf("%d", num)
			}
		}
		fmt.Printf("\n")
	}
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