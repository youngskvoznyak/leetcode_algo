package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	record := make(map[*Node]*Node)
	dfs(node, record)

	return record[node]

}

func dfs(node *Node, record map[*Node]*Node) {
	if node == nil {
		return
	}
	newNode := new(Node)
	newNode.Val = node.Val

	record[node] = newNode

	for _, child := range node.Neighbors {
		if _, exist := record[child]; !exist {
			dfs(child, record)
		}
		newNode.Neighbors = append(newNode.Neighbors, record[child])
	}
}
