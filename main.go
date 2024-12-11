package main

func main() {
	nodes := []string{"A", "B", "C", "D", "E"}
	nodeMap := make(map[string]int)
	for i, node := range nodes {
		nodeMap[node] = i
	}

	matrix := make([][]int, len(nodes))
	for i := range matrix {
		matrix[i] = make([]int, len(nodes))
	}

	AdjMatGraph := AdjMatGraph{Nodes: nodes, Matrix: matrix, NodeMap: nodeMap}

	AdjMatGraph.AddEdge("A", "B")
	AdjMatGraph.AddEdge("A", "C")
	AdjMatGraph.AddEdge("B", "D")
	AdjMatGraph.AddEdge("C", "D")
	AdjMatGraph.AddEdge("D", "E")

	AdjMatGraph.PrintMatrix()

	AdjListGraph := AdjListGraph{
		AdjList: make(map[string][]string),
	}

	AdjListGraph.AddEdge("A", "B")
	AdjListGraph.AddEdge("A", "C")
	AdjListGraph.AddEdge("B", "D")
	AdjListGraph.AddEdge("C", "D")
	AdjListGraph.AddEdge("D", "E")

	AdjListGraph.PrintAdjListGraph()
}
