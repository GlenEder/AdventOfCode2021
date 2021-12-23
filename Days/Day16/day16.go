package Day16

import (
	"fmt"
	"src/Utils"
)

func Run() {
	Utils.PrintDay(16)

	//get input
	input := Utils.ReadInputAsString("Days/Day16/input.txt")
	input = Utils.RemoveWhiteSpace(input)
	//fmt.Printf("Decoding %s\n", input)
	//convert from hex to binary
	binary := Utils.HexToBinary(input)
	//fmt.Printf("%s\n", binary)

	part1, part2, _ := eval(binary, 0)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func eval(binary string, pos int) (int, int, int) {

	if Utils.BinaryToInt(binary[pos:]) == 0 {
		return 0, 0, pos
	}

	//Get version number
	version := Utils.BinaryToInt(binary[pos : pos+3])
	//add version for part 1
	part1 := version
	//get type id
	typeId := Utils.BinaryToInt(binary[pos+3 : pos+6])
	pos += 6

	//handle literal
	if typeId == 4 {
		val, dist := evalLiteral(binary[pos:])
		pos += dist
		//return evaled literal data
		return part1, val, pos
	} else {
		//Get length id
		lenID := Utils.BinaryToInt(binary[pos : pos+1])
		pos++

		var subVals []int

		//Handle 15 bit version
		if lenID == 0 {
			subPackLen := Utils.BinaryToInt(binary[pos : pos+15])
			pos += 15
			start := pos
			for {
				p1Add, subval, newPos := eval(binary, pos)
				part1 += p1Add
				subVals = append(subVals, subval)
				pos = newPos
				if pos-start >= subPackLen {
					break
				}
			}

		} else {
			//Handle 11 bit version
			numSubpacks := Utils.BinaryToInt(binary[pos : pos+11])
			pos += 11
			for i := 0; i < numSubpacks; i++ {
				p1Add, subval, newPos := eval(binary, pos)
				part1 += p1Add
				subVals = append(subVals, subval)
				pos = newPos

			}
		}

		//Eval on type
		eval := evalType(typeId, subVals)
		//fmt.Printf("Evaled type %d on %v = %d\n", typeId, subVals, eval)
		return part1, eval, pos
	}
}

func evalLiteral(binary string) (int, int) {
	//Literal value
	decoded := ""
	reading := true
	pos := 0
	for reading {
		if binary[pos:pos+1] == "0" {
			reading = false
		}
		decoded += binary[pos+1 : pos+5]
		pos += 5
	}

	toReturn := Utils.BinaryToInt(decoded)
	//Utils.PrintWithColorf(Utils.Green, "decoded: %d\n", toReturn)
	return toReturn, pos
}

func evalType(typeId int, toCombine []int) int {

	switch typeId {
	case 0:
		return Utils.SumOfIntSlice(toCombine)
	case 1:
		return Utils.ProductOfIntSlice(toCombine)
	case 2:
		min, _ := Utils.MinOfIntSlice(toCombine)
		return min
	case 3:
		max, _ := Utils.MaxOfIntSlice(toCombine)
		return max
	case 5:
		if toCombine[0] > toCombine[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if toCombine[0] < toCombine[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if toCombine[0] == toCombine[1] {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
