package main

import "fmt"

type AdjListGraph struct {
	AdjList map[string][]string
}

func (alg *AdjListGraph) AddEdge(node, neighbor string) {
	alg.AdjList[node] = append(alg.AdjList[node], neighbor)
}

func (alg *AdjListGraph) PrintAdjListGraph() {
	for node, neighbors := range alg.AdjList {
		fmt.Printf("%s -> %v\n", node, neighbors)
	}
}
