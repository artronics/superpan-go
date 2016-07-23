package device

import "net"

type Device struct {
	Receiver
	MACAddress MACAddress
	DeviceType DeviceType
}

type MACAddress net.HardwareAddr

type Receiver struct {
}

type DeviceType int

const (
	FFD DeviceType = iota
	RFD
)
