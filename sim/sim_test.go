package sim

import (
	"github.com/artronics/superpan/node"
	"testing"
)

func TestSim_New(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatalf("sim instance must not be nil")
	}
}

func TestSim_AddNode_should_add_node_to_sim(t *testing.T) {
	sim := New()
	sim.AddNode(node.PanNode{})

	if len(sim.nodes) != 1 {
		t.Fatalf("expect 1 node but got %d", len(sim.nodes))
	}
}
