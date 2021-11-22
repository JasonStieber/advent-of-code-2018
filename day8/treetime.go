package main

import "log"

type node struct {
	metaData []int
	children []*node
}

func main() {
	head, _ := buildTree(0, input)
	total := countData(head)
	log.Printf("The answer to part A the sum of all the metadata is: %v \n", total)
}

func countData(tree node) int {
	total := 0
	for i := 0; i < len(tree.children); i++ {
		total += countData(*tree.children[i])
	}
	for j := 0; j < len(tree.metaData); j++ {
		total += tree.metaData[j]
	}
	return total
}

func buildTree(loc int, data []int) (node, int) {
	var n node
	nLoc := loc + 2
	numKids := data[loc]
	numData := data[loc+1]
	for i := 0; i < numKids; i++ {
		var child node
		child, nLoc = buildTree(nLoc, data)
		n.children = append(n.children, &child)
	}
	for j := 0; j < numData; j++ {
		n.metaData = append(n.metaData, data[nLoc+j])
	}
	return n, nLoc + numData
}
