package main

import "fmt"

type Node struct {
	key      string
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

func	queueInit(tree Node) (queue[] Node) {
	
	if tree.child != nil {
		queue = append(queue, tree)
		for _, child := range tree.child {
			queue = append(queue, queueInit(child)...)
		}
	} else {
		queue = append(queue, tree)
	}
	return 
}

func nodeIsValid(tree Node) (bool) {
	if tree.key[len(tree.key) - 1] == 'e' {
		return true
	} 
	return false
}

func main() {
	l_1_1, l_1_2 := Node{key: "Biba"}, Node{key: "Boba"}
	l_2_1, l_2_2 := Node{key: "Vasya"}, Node{key: "Petya"}
	l_1, l_2 := Node{key: "Alexa"}, Node{key:"Alice"}
	l := Node{key: "MeIdiot"}
	l_1.child = append(l_1.child, l_1_1, l_1_2)
	l_2.child = append(l_2.child, l_2_1, l_2_2)
	l.child = append(l.child, l_1, l_2)

	queue := queueInit(l)
	for _, q := range queue {
		if nodeIsValid(q) {
			fmt.Println(q.key)
		}
	}
}
