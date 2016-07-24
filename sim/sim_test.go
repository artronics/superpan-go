package sim

import (
	"testing"
	"github.com/artronics/superpan/device"
)

func TestSim_AddNodes(t *testing.T) {
	n1:=PanNode{device:&device.Device{MACAddress:[]byte{0,1}}}
	n2:=PanNode{device:&device.Device{MACAddress:[]byte{0,2}}}
	//same as n2
	n3:=PanNode{device:&device.Device{MACAddress:[]byte{0,2}}}
	Sim.addNodes(n1,n2,n3)
	if l := len(Sim.nodes); l != 2 {
		t.Errorf("expected 2 items but got %d",l)
	}
}
