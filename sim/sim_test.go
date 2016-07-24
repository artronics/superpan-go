package sim

import "testing"

var id=0
func TestSim_AddNodes(t *testing.T) {
	n1:=node{}
	n2:=node{}
	Sim.addNodes(n1,n2)
	if l := len(Sim.nodes); l != 2 {
		t.Errorf("expected 2 items but got %d",l)
	}
}
type node struct {}

func (n node)ID() string {
	id++
	return string(id)
}
