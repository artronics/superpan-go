package sim

import (
	"github.com/artronics/superpan/device"
)

var Sim sim = sim{
	nodes:make(map[string]Node),
}

type sim struct {
	nodes map[string]Node
}

func (s *sim)Start() {

}

type Node interface {
	ID() string
	Device() *device.Device
}

type PanNode struct {
	device *device.Device
}

func (n PanNode)ID() string {
	return string(n.device.MACAddress)
}

func (n PanNode)Device() *device.Device {
	return n.device
}

func (s *sim)addNodes(nodes... Node) {
	for _,node := range nodes {
		Sim.nodes[node.ID()]=node
	}
}
