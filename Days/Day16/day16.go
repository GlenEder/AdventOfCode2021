package Day16

import (
	"fmt"
	"src/Utils"
)

func Run() {
	Utils.PrintDay(16)

	//get input
	input := Utils.ReadInputAsString("Days/Day16/testInput5.txt")
	input = Utils.RemoveWhiteSpace(input)
	//fmt.Printf("Decoding %s\n", input)
	//convert from hex to binary
	binary := Utils.HexToBinary(input)
	//fmt.Printf("%s\n", binary)

	part2total, _, part1total := evalPacket(binary, 0)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1total, part2total[0])
}

func evalPacket(binary string, depth int) ([]int, int, int) {
	Utils.PrintRecursionLogColorf(Utils.Yellow, depth, "Starting eval of: %s\n", binary)
	toReturn := []int{}
	part1Return := 0
	i := 0
	for i+6 < len(binary) {
		//check for only 0's left
		if Utils.BinaryToInt(binary[i:]) == 0 {
			i = -1 //signal packet is done
			break
		}

		//Split into sections
		version := Utils.BinaryToInt(binary[i : i+3])
		typeId := Utils.BinaryToInt(binary[i+3 : i+6])
		//calc part 1
		part1Return += version

		Utils.PrintRecursionLogf(depth, "version: %d\ttype id:%d\tdata:%s\n", version, typeId, binary[i+6:])

		if typeId == 4 {
			val, len := evalLiteral(binary[i+6:], depth)
			toReturn = append(toReturn, val)
			i += len + 6
		} else {
			var subNumbers []int
			//Handle operator
			lengthType := binary[i+6 : i+7]
			if lengthType == "1" {
				//the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
				frontOfPack := i + 7
				endOfPack := frontOfPack + 11
				numPacks := Utils.BinaryToInt(binary[frontOfPack:endOfPack])
				Utils.PrintRecursionLogf(depth, "Number of subpackets: %d\n", numPacks)
				for j := 0; j < numPacks; j++ {
					toAdd, len, part1 := evalPacket(binary[endOfPack:], depth+1)
					subNumbers = append(subNumbers, toAdd...)
					endOfPack += len
					part1Return += part1
				}
				i = endOfPack
			} else {
				//the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet.
				//Utils.PrintWithColorf(Utils.Yellow, "Defined supacket length found\n")
				subpackLen := Utils.BinaryToInt(binary[i+7 : i+22])
				Utils.PrintRecursionLogf(depth, "subpacket length: %d\n", subpackLen)
				toAdd, _, part1 := evalPacket(binary[i+22:i+22+subpackLen], depth+1)
				subNumbers = append(subNumbers, toAdd...)
				part1Return += part1
				//TODO calc i advancement
				i += 22 + subpackLen
			}

			//perform eval
			eval := evalType(typeId, subNumbers, depth)
			toReturn = append(toReturn, eval)
		}
	}
	Utils.PrintRecursionLogf(depth, "Returning: %v\n", toReturn)
	return toReturn, i, part1Return
}

func evalLiteral(binary string, depth int) (int, int) {
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
	Utils.PrintRecursionLogColorf(Utils.Cyan, depth, "decoded: %s, %d\n", decoded, toReturn)
	return toReturn, pos
}

func evalType(typeId int, toCombine []int, depth int) int {
	Utils.PrintRecursionLogf(depth, "Evaling type: %d, w/ %v\n", typeId, toCombine)
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
