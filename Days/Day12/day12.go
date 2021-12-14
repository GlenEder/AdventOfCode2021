package Day12

import (
	"fmt"
	"src/Utils"
	"strings"
)

type node struct {
	name string
	connections []string
}

func Run() {
	Utils.PrintDay(12)
	//get input
	input := Utils.ReadInputToSlice("Days/Day12/input.txt")
	//node mapping
	var nodes []node

	//load nodes
	for _, line := range input {
		parts := strings.Split(line, "-")
		//Add A -> B
		index := nodeExists(parts[0], nodes)
		if index < 0 {
			//Add new node
			nodes = append(nodes, node{
				name: parts[0],
				connections: []string{parts[1]},
			})
		} else {
			//Add new connection
			nodes[index].connections = append(nodes[index].connections, parts[1])
		}

		//Add B -> A
		index = nodeExists(parts[1], nodes)
		if index < 0 {
			//Add new node
			nodes = append(nodes, node{
				name: parts[1],
				connections: []string{parts[0]},
			})
		} else {
			//Add new connection
			nodes[index].connections = append(nodes[index].connections, parts[0])
		}
	}
	//fmt.Printf("%v\n", nodes)
	//Get start node
	var startNode node
	for _, n := range nodes {
		if n.name == "start" {
			startNode = n
			break
		}
	}

	numPaths := explore(startNode, []string{}, nodes)
	part2Paths := exploreP2(startNode, []string{},  nodes, 0)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", numPaths, part2Paths)
}

func explore(start node, path []string, nodes []node) int {
	numPathsFound := 0

	//go through all connections available
	for _, c := range start.connections {
		//fmt.Printf("\tChecking connection: %s\tPath: %v\n", c, path)
		//If visited small cave don't count path
		if Utils.IsLowerCase(c) && Utils.SliceContainsString(path, c) {
			//Utils.PrintWithColorf(Utils.Red, "\t\tCave already visited: %s\n", c)
		} else if c == "end" {
			//Utils.PrintWithColorf(Utils.Green, "\t\tPath found: %v + end\n", path)
			numPathsFound++
		} else {
			//Explode on node
			index := nodeExists(c, nodes)
			numPathsFound += explore(nodes[index], append(path, start.name), nodes)
		}
	}

	return numPathsFound
}

func exploreP2(start node, path []string, nodes []node, depth int) int {
	numPathsFound := 0
	//go through all connections available
	newPath := append(path, start.name)
	canDoubleVisit := canDouble(newPath)
	//Utils.PrintRecursionLogf(depth, "Exploring node %s, connections: %v\tPath: %v\n", start.name, start.connections, newPath)
	for _, c := range start.connections {
		visits := Utils.NumStringMatches(newPath, c)
		//Utils.PrintRecursionLogf(depth,"Checking connection: %s\tNum Visits: %d\tCan Double: %t\n", c, visits, canDoubleVisit)
		if c == "start" && visits > 0 {
			//Utils.PrintRecursionLogf(depth, "Start already visited\n")
		} else if Utils.IsLowerCase(c) && visits > 0 && !canDoubleVisit{
			//Utils.PrintRecursionLogf(depth, "Cave %s already visited %d times\n", c, visits)
		} else if c == "end" {
			//Utils.PrintWithColorf(Utils.Green, printFoundPath(newPath))
			numPathsFound++
		} else {
			//Explode on node
			index := nodeExists(c, nodes)
			numPathsFound += exploreP2(nodes[index], newPath, nodes, depth + 1)
		}
	}

	return numPathsFound
}

func canDouble(path []string) bool {
	double := make(map[string]int)
	for _, n := range path {
		if Utils.IsLowerCase(n) {
			if _, exists := double[n]; exists {
				return false
			} else {
				double[n] = 1
			}
		}
	}
	return true
}

func printFoundPath(s []string) string {
	toprint := ""
	for _, n := range s {
		toprint += n + ", "
	}
	toprint += "end\n"
	return toprint
}

//Returns index of node in list if exists in list
//@param name -- name to look for
//@param nodes -- node slice to look in
//@return index of node if exists, -1 otherwise
func nodeExists(name string, nodes []node) int {
	for i, n := range nodes {
		if n.name == name {
			return i
		}
	}
	return -1
}