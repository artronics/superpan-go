package sim

import "github.com/artronics/superpan/node"

type Sim struct {
	nodes     []node.Node
	subscribe chan node.Node
}

func New() *Sim {
	s := &Sim{subscribe: make(chan node.Node)}
	return s
}

func (s *Sim) sub() {

}

func (s *Sim) AddNode(node node.Node) {
	s.nodes = append(s.nodes, node)
}
