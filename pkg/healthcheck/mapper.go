package healthcheck

import (
	"fmt"
	"log"
)

type Node struct {
	name, endpoint string
	postNodes      []string
}

func (n *Node) Equals(other *Node) bool {
	return n.name == other.name
}

func NewNode(name string, endpoint string, postNodes []string) *Node {
	return &Node{name: name, endpoint: endpoint, postNodes: postNodes}
}

func Run() []Node {
	nodes := initNodes()
	DisplayNodes(nodes)
	return nodes
}

func DisplayNodes(nodeList []Node) {
	mappedNodes := make(map[string][]*Node)
	fmt.Println(nodeList)
	for _, node := range nodeList {
		var postNodes []*Node

		for _, postNodeName := range node.postNodes {
			fmt.Println(postNodeName)

			foundNode := findNodeByName(nodeList, postNodeName)
			postNodes = append(postNodes, foundNode)
		}

		//fmt.Printf("v: %v\ts: %s\tT: %T\n", node, node, node)
		mappedNodes[node.name] = postNodes
	}

	for k, v := range mappedNodes {
		fmt.Println(k, "value is", v)
	}

}

func findNodeByName(nodeList []Node, name string) *Node {
	for _, n := range nodeList {
		if n.name == name {
			return &n
		}
	}
	log.Fatalf("Could not found node %s\n", name)
	return nil
}

func initNodes() []Node {
	nodeSlice := []Node{
		*NewNode("name1", "endpoint1", nil),
		*NewNode("name2", "endpoint2", []string{"name1"}),
		//*NewNode("name3", "endpoint3", nil),
	}
	return nodeSlice
}
