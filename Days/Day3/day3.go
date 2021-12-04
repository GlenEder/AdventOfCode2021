package Day3

import (
	"fmt"
	"math"
	"src/Utils"
)

//Holds count for
type bitCount struct {
	num1s int
	num0s int
}

func Run() {
	fmt.Printf("===Day 3===\n")
	//Read input
	input := Utils.ReadInputToSlice("Days/Day3/input.txt")
	//Find number of bits
	numBits := len(input[0])
	//create slice of bitCounters
	bits := make([]bitCount, numBits)
	//loop through input
	for _, inputLine := range input {
		//loop through chars in input line
		for i, c := range inputLine {
			if c == '1' {
				bits[i].num1s++
			} else {
				bits[i].num0s++
			}
		}
	}


	var gammaRate, epsilonRate  int
	//Calc nums
	for i, bit := range bits {
		pos := float64(numBits - i - 1)
		if bit.num1s > bit.num0s {
			gammaRate += int(math.Exp2(pos))
		} else {
			epsilonRate += int(math.Exp2(pos))
		}
	}

	fmt.Printf("Part 1: %d\n", gammaRate * epsilonRate)

	oxygenReading := 0
	co2Rating := 0

	//Handle oxygen first
	var toLoopOn, oneStrings, zeroStrings []string
	toLoopOn = input
	bitToTest := 0

	for bitToTest < numBits {
		//loop through input left
		for _, inputLine := range toLoopOn {
			if inputLine[bitToTest] == '1' {
				oneStrings = append(oneStrings, inputLine)
			} else {
				zeroStrings = append(zeroStrings, inputLine)
			}
		}

		//Check for which bit is more popular
		if len(zeroStrings) > len(oneStrings) {
			toLoopOn = zeroStrings
		} else {
			toLoopOn = oneStrings
		}

		//Check for num found
		if len(toLoopOn) == 1 {
			oxygenReading = int(Utils.BinaryStringToFloat(toLoopOn[0]))
			break
		}
		//move to next bit and reset tracking arrays
		bitToTest++
		oneStrings = nil
		zeroStrings = nil
	}

	// Handle co2 reading
	toLoopOn = input
	oneStrings = nil
	zeroStrings = nil
	bitToTest = 0

	for bitToTest < numBits {
		//loop through input left
		for _, inputLine := range toLoopOn {
			if inputLine[bitToTest] == '1' {
				oneStrings = append(oneStrings, inputLine)
			} else {
				zeroStrings = append(zeroStrings, inputLine)
			}
		}

		//Check for which bit is less popular
		if len(zeroStrings) > len(oneStrings) {
			toLoopOn = oneStrings
		} else {
			toLoopOn = zeroStrings
		}

		//Check for num found
		if len(toLoopOn) == 1 {
			co2Rating = int(Utils.BinaryStringToFloat(toLoopOn[0]))
			break
		}
		//move to next bit and reset tracking arrays
		bitToTest++
		oneStrings = nil
		zeroStrings = nil
	}

	fmt.Printf("Part 2: %d\n", oxygenReading * co2Rating)


}