package util

import (
	"container/list"
	"sort"
)

type Node struct {
	Id      string
	IsSmall bool
	Friends map[string]*Node
}

type ValueNode struct {
	Id      string
	Value   int
	Friends map[string]ValueNode
}

func DFS(n *Node, array []string, endId string) []string {

	array = append(array, n.Id)
	if n.Id == endId {
		return array
	}

	for _, child := range n.Friends {
		array = DFS(child, array, endId)
	}

	return array
}

func DFSValue(n *ValueNode, array []string, endId string) []string {

	array = append(array, n.Id)
	if n.Id == endId {
		return array
	}

	for _, child := range n.Friends {

		array = DFSValue(&child, array, endId)

	}

	return array
}

func BFS(n *Node) []*Node {

	//track the visited nodes
	visited := make(map[string]*Node)
	// queue of the nodes to visit
	queue := list.New()
	queue.PushBack(n)
	// add the root node to the map of the visited nodes
	visited[n.Id] = n

	for queue.Len() > 0 {
		qnode := queue.Front()
		// iterate through all of its friends
		// mark the visited nodes; enqueue the non-visted
		for id, node := range qnode.Value.(*Node).Friends {
			if _, ok := visited[id]; !ok {
				visited[id] = node
				queue.PushBack(node)
			}
		}
		queue.Remove(qnode)
	}

	nodes := make([]*Node, 0)
	// collecte all the nodes into slice
	for _, node := range visited {
		nodes = append(nodes, node)
	}

	return nodes
}

func AllPathSearch(currentNode *Node, currentPath []*Node, endId string) [][]*Node {

	if currentNode.Id == endId {
		currentPath = append(currentPath, currentNode)
		return [][]*Node{currentPath}
	}

	allPaths := [][]*Node{}
	currentPath = append(currentPath, currentNode)
	for _, friendNode := range currentNode.Friends {
		newPath := append([]*Node{}, currentPath...)
		allPaths = append(allPaths, AllPathSearch(friendNode, newPath, endId)...)
	}

	return allPaths
}

func AllPathSearchValueNodes(currentNode *ValueNode, currentPath []ValueNode, endId string, iterations int) [][]ValueNode {
	currentNodeCopy := ValueNode{currentNode.Id, currentNode.Value, currentNode.Friends}

	if currentNode.Id == endId {
		currentPath = append(currentPath, currentNodeCopy)

		return [][]ValueNode{currentPath}
	}

	//Don't go backwards
	currentPathStr := []string{}
	for _, c := range currentPath {
		currentPathStr = append(currentPathStr, c.Id)
		if currentNode.Id == c.Id {
			return [][]ValueNode{}
		}
	}

	allPaths := [][]ValueNode{}

	currentPath = append(currentPath, currentNodeCopy)
	sortedFriendNodes := SortValueNodesAsc(currentNodeCopy.Friends)

	for _, friendNode := range sortedFriendNodes {
		if InStringSlice(currentPathStr, friendNode.Id) {
			delete(currentNodeCopy.Friends, friendNode.Id)
			continue
		}
		newPath := append([]ValueNode{}, currentPath...)
		iterations++
		if iterations > 100 {
			panic("STAAAP")
		}

		allPaths = append(allPaths, AllPathSearchValueNodes(&friendNode, newPath, endId, iterations)...)
	}

	return allPaths
}

func SortValueNodesAsc(friendNodes map[string]ValueNode) []ValueNode {
	friendNodeValues := []int{}
	for _, friendNode := range friendNodes {
		friendNodeValues = append(friendNodeValues, friendNode.Value)
	}

	sort.Ints(friendNodeValues)
	result := make([]ValueNode, len(friendNodeValues))

	for _, friendNode := range friendNodes {
		for i, friendNodeVal := range friendNodeValues {
			if friendNode.Value == friendNodeVal {
				result[i] = friendNode
				break
			}
		}
	}

	return result
}

func AllPathSearchWithNodePredicate(currentNode *Node, currentPath []*Node, endId string, nodePredicateFn func(*Node, []*Node) bool) [][]*Node {

	if currentNode.Id == endId {
		currentPath = append(currentPath, currentNode)
		return [][]*Node{currentPath}
	}

	if nodePredicateFn(currentNode, currentPath) {
		return [][]*Node{}
	}

	allPaths := [][]*Node{}
	currentPath = append(currentPath, currentNode)
	for _, friendNode := range currentNode.Friends {
		newPath := append([]*Node{}, currentPath...)
		allPaths = append(allPaths, AllPathSearchWithNodePredicate(friendNode, newPath, endId, nodePredicateFn)...)
	}

	return allPaths
}

func AllPathSearchWithPathPredicates(currentNode *Node, currentPath []*Node, endId string, pathPredicateFn func(*Node, []*Node, bool) (bool, bool), pathPredicate bool) [][]*Node {

	if currentNode.Id == endId {
		currentPath = append(currentPath, currentNode)
		return [][]*Node{currentPath}
	}

	newPath, visitedSmall := pathPredicateFn(currentNode, currentPath, pathPredicate)
	if newPath {
		return [][]*Node{}
	}

	allPaths := [][]*Node{}
	currentPath = append(currentPath, currentNode)
	for _, friendNode := range currentNode.Friends {
		newPath := append([]*Node{}, currentPath...)
		allPaths = append(allPaths, AllPathSearchWithPathPredicates(friendNode, newPath, endId, pathPredicateFn, visitedSmall)...)
	}

	return allPaths
}
