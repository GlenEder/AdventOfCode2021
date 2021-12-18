package main

import (
	"src/Days/Day1"
	"src/Days/Day10"
	"src/Days/Day11"
	"src/Days/Day12"
	"src/Days/Day13"
	"src/Days/Day14"
	"src/Days/Day15"
	"src/Days/Day2"
	"src/Days/Day3"
	"src/Days/Day4"
	"src/Days/Day5"
	"src/Days/Day6"
	"src/Days/Day7"
	"src/Days/Day8"
	"src/Days/Day9"
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
	Day7.Run()
	Day8.Run()
	Day9.Run()
	Day10.Run()
	Day11.Run()
	Day12.Run()
	Day13.Run()
	Day14.Run()
	Day15.Run()
	Utils.PrintWithColor(Utils.Yellow, "\nTime to compute: " + time.Since(now).String() + "\n")
}
