package main

import (
	"github.com/artronics/superpan/app"
	"github.com/artronics/superpan/ieee8021504"
	"github.com/artronics/superpan/ieee8021504/primitive"
	"github.com/artronics/superpan/node"
	"log"
)

func main() {
	i := ieee8021504.IEEE8021504{}
	n := node.Node{App: &app.App{IEEE8021504: &i}, Device: &node.Device{DeviceType: node.FFD}}
	log.Println(n)
	i.MLME.Request(primitive.ScanRequest{})
}
