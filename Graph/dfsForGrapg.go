/*
https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
Depth First Search or DFS for a Graph
*/

package main

import (
	"container/list"
	"fmt"
)

// NodeAttr stores nodes and edges
type NodeAttr struct {
	edges map[*Node][]*Node
}

// Node stores values
type Node struct {
	value   int
	visited bool
}

// addEdges connect two nodes with edges
func (na *NodeAttr) addEdges(n1, n2 *Node) {
	if na.edges == nil {
		na.edges = make(map[*Node][]*Node)
	}

	na.edges[n1] = append(na.edges[n1], n2)
	na.edges[n2] = append(na.edges[n2], n1)
}

// DFS depth first search
func (na *NodeAttr) DFS(n *Node) {
	var v []*Node
	var ok bool

	if v, ok = na.edges[n]; ok {
		if !n.visited {
			fmt.Println(n.value)
			n.visited = true
		}
	}

	for _, v1 := range v {
		if !v1.visited {
			na.DFS(v1)
		}
	}
}

// DFSStack depth first search using stack
func (na *NodeAttr) DFSStack(n *Node) {

	if _, ok := na.edges[n]; !ok {
		return
	}

	ll := list.New()
	ll.PushBack(n)

	for ll.Len() != 0 {
		var v []*Node
		var ok bool

		elem := ll.Front()

		if v, ok = na.edges[elem.Value.(*Node)]; ok {
			if !elem.Value.(*Node).visited {
				fmt.Println(elem.Value.(*Node).value)
				elem.Value.(*Node).visited = true
			}
		}

		for _, v1 := range v {
			if !v1.visited {
				ll.PushFront(v1)
			}
		}

		ll.Remove(elem)
	}
}

func main() {
	nodes := NodeAttr{}
	a0 := &Node{value: 0}
	a1 := &Node{value: 1}
	a2 := &Node{value: 2}
	a3 := &Node{value: 3}

	nodes.addEdges(a0, a1)
	nodes.addEdges(a0, a2)
	nodes.addEdges(a1, a2)
	nodes.addEdges(a2, a0)
	nodes.addEdges(a2, a3)
	nodes.addEdges(a3, a3)

	fmt.Println("DFS starting from node 2 using recursive:")
	nodes.DFS(a2)

	nodes1 := NodeAttr{}
	aa0 := &Node{value: 0}
	aa1 := &Node{value: 1}
	aa2 := &Node{value: 2}
	aa3 := &Node{value: 3}
	aa4 := &Node{value: 4}

	nodes1.addEdges(aa1, aa0)
	nodes1.addEdges(aa0, aa2)
	nodes1.addEdges(aa2, aa1)
	nodes1.addEdges(aa0, aa3)
	nodes1.addEdges(aa1, aa4)

	fmt.Println("DFS starting from node 0 using stack:")
	nodes1.DFSStack(aa0)
}

/*
DFS starting from node 2 using recursive:
2
0
1
3
DFS starting from node 0 using stack:
0
3
2
1
4
*/
