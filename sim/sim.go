package sim

import "fmt"

var Sim sim = sim{
	nodes:make(map[string]Node,1),
}

type sim struct {
	nodes map[string]Node
}

type Node interface {
	ID() string
}

func (s *sim)addNodes(nodes... Node) {
	fmt.Printf("%v",Sim.nodes)
	for _,node := range nodes {
		Sim.nodes[node.ID()]=node
	}
}
