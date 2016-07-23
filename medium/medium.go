package medium

import (
	"github.com/artronics/superpan/device"
	"net"
)

type Medium struct {
	devices map[string]*device.Device
}

func (m *Medium) AddDevice(devices ...*device.Device) {
	if m.devices == nil {
		m.devices = make(map[string]*device.Device, len(devices))
	}
	for _, device := range devices {
		mac := net.HardwareAddr(device.MACAddress).String()
		m.devices[mac] = device
	}
}

func (m *Medium) RemoveDevice(devices ...*device.Device) {
	for _, device := range devices {
		mac := net.HardwareAddr(device.MACAddress).String()
		delete(m.devices, mac)
	}
}
