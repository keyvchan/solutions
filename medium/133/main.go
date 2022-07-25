package main

import "fmt"

func main() {
	// test case
	// create 4 nodes
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	// create graph
	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)

	// clone graph
	new_node := cloneGraph(node1)
	for _, neighbor := range new_node.Neighbors {
		fmt.Println(neighbor.Val)
		for _, neighbor := range neighbor.Neighbors {
			fmt.Println(" ", neighbor.Val)
			for _, neighbor := range neighbor.Neighbors {
				fmt.Println("  ", neighbor.Val)
			}
		}
	}

	// create a empty graph
	node1 = &Node{Val: 1}
	new_node = cloneGraph(node1)
	fmt.Println(new_node)

	new_node = cloneGraph(nil)
	fmt.Println(new_node)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(node *Node, new_graph *Node, visited_nodes map[*Node]*Node) *Node {
	// if node is visited, return
	if _, ok := visited_nodes[node]; ok {
		return visited_nodes[node]
	}
	if node == nil {
		return visited_nodes[node]
	}
	visited_nodes[node] = new_graph
	// copy the current node
	new_graph.Val = node.Val
	// copy the neighbors
	for _, neighbor := range node.Neighbors {
		// create a new neighbor
		new_neighbor := &Node{}
		new_node := dfs(neighbor, new_neighbor, visited_nodes)
		// add to the neighbors
		new_graph.Neighbors = append(new_graph.Neighbors, new_node)
	}
	return new_graph
}

func cloneGraph(node *Node) *Node {
	// create a new graph
	new_root := &Node{}
	visited_nodes := map[*Node]*Node{}
	// bfs
	new_root = dfs(node, new_root, visited_nodes)
	return new_root
}
