package main

import "fmt"

type AdjMatGraph struct {
	Nodes   []string
	Matrix  [][]int
	NodeMap map[string]int
}

func (amg *AdjMatGraph) AddEdge(node1, node2 string) {
	i, j := g.NodeMap[node1], amg.NodeMap[node2]
	amg.Matrix[i][j] = 1
}

func (amg *AdjMatGraph) PrintAdjMatGraph() {
	fmt.Printf("   %v\n", amg.Nodes)
	for i, row := range amg.Matrix {
		fmt.Printf("%s: %v\n", amg.Nodes[i], row)
	}
}
