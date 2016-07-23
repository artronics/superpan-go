package medium

import (
	"github.com/artronics/superpan/device"
	"testing"
)

func TestMedium_AddDevice(t *testing.T) {
	m := Medium{}
	d1 := &device.Device{MACAddress: []byte{0, 1, 3, 4}}
	d2 := &device.Device{MACAddress: []byte{0, 1, 3, 5}}
	//same address as d2
	d3 := &device.Device{MACAddress: []byte{0, 1, 3, 5}}
	m.AddDevice(d1, d2, d3)
	if c := len(m.devices); c != 2 {
		t.Errorf("expected 2 devices but got %d", c)
	}
}

func TestMedium_RemoveDevice(t *testing.T) {
	m := Medium{}
	d1 := &device.Device{MACAddress: []byte{0, 1, 3, 4}}
	d2 := &device.Device{MACAddress: []byte{0, 1, 3, 5}}
	//same address as d2
	d3 := &device.Device{MACAddress: []byte{0, 1, 3, 5}}
	m.AddDevice(d1, d2, d3)
	m.RemoveDevice(d3)
	if c := len(m.devices); c != 1 {
		t.Errorf("expected 1 devices but got %d", c)
	}
}
