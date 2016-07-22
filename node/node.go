package node

import (
	"github.com/artronics/superpan/app"
)

type Node struct {
	App    *app.App
	Device *Device
}

func (p *Node) Start() {

}
