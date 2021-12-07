package main

import (
	"src/Days/Day1"
	"src/Days/Day2"
	"src/Days/Day3"
	"src/Days/Day4"
	"src/Days/Day5"
	"src/Days/Day6"
	"src/Utils"
	"time"
)

func main() {
	now := time.Now()
	Day1.Run()
	Day2.Run()
	Day3.Run()
	Day4.Run()
	Day5.Run()
	Day6.Run()
	Utils.PrintWithColor(Utils.Yellow, "\nTime to compute: " + time.Since(now).String() + "\n")
}
