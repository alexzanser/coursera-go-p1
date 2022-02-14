package main

import "fmt"

type Node struct {
	key   string
	child []Node
}

func showTree(tree Node) {
	if tree.child != nil {
		fmt.Println(tree.key)
		for _, child := range tree.child {
			showTree(child)
		}
	} else {
		fmt.Println(tree.key)
	}
}

func queueInit(tree Node) (queue []Node) {

	queue = append(queue, tree)
	if tree.child != nil {
		for _, child := range tree.child {
			queue = append(queue, child)
		}
		for _, child := range tree.child {
			for _, sub_child := range child.child {
				queue = append(queue, queueInit(sub_child)...)
			}
		}
	}
	return
}

func nodeIsValid(tree Node) bool {
	if tree.key[len(tree.key)-1] == 'z' {
		return true
	}
	return false
}

func main() {
	l_1_1, l_1_2 := Node{key: "Biba"}, Node{key: "Boba"}
	l_2_1, l_2_2 := Node{key: "Vasya"}, Node{key: "Petyz"}
	l_1, l_2 := Node{key: "Alexa", child : []Node {l_1_1, l_1_2}}, 
				Node{key: "Alice", child : []Node {l_2_1, l_2_2}}
	l := Node{key: "MeIdiot", child : []Node {l_1, l_2}}

	queue := queueInit(l)
	for _, q := range queue {
		if nodeIsValid(q) {
			fmt.Println(q.key)
			break
		}
	}
}
