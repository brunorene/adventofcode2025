package day11

import (
	"adventofcode2025/common"
	"fmt"
	"slices"
	"strings"
)

type Node struct {
	name     string
	children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		name:     name,
		children: make([]*Node, 0),
	}
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) paths(endName string, mandatoryNodes []string, visited map[string][4]int) [4]int {
	if cached, exists := visited[n.name]; exists {
		return cached
	}

	if n.name == endName {
		return [4]int{1, 0, 0, 0}
	}
	var result [4]int
	idx := slices.Index(mandatoryNodes, n.name)
	for _, child := range n.children {
		counters := child.paths(endName, mandatoryNodes, visited)
		if idx < 0 { // it is not mandatory node or we already passed by both
			for i := range 4 {
				result[i] += counters[i]
			}
			continue
		}

		if counters[1<<idx^3] > 0 { // it is a mandatory node and we passed by the other before
			result[3] += counters[1<<idx^3]
		}

		if counters[0] > 0 {
			result[1<<idx] += counters[0]
		}
	}

	visited[n.name] = result

	return result
}

func Solve(input string, mandatoryNodes []string, start string) int {
	stop := common.Timer("solution")
	defer stop()

	fmt.Println("Solving for input:", input, "Mandatory nodes:", mandatoryNodes, "Start:", start)

	nodes := make(map[string]*Node)

	for line := range common.ReadInput(input).ReadLines {
		parts := strings.Split(line, ": ")
		parent, exists := nodes[parts[0]]
		if !exists {
			parent = NewNode(parts[0])
			nodes[parent.name] = parent
		}

		children := strings.SplitSeq(parts[1], " ")
		for childName := range children {
			child, exists := nodes[childName]
			if !exists {
				child = NewNode(childName)
				nodes[childName] = child
			}
			parent.AddChild(child)
		}
	}

	result := nodes[start].paths("out", mandatoryNodes, map[string][4]int{})

	if len(mandatoryNodes) == 0 {
		return result[0]
	}

	return result[3]
}
