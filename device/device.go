package device

import (
	"net"
)

type Device struct {
	receiver
	transmitter
	medium chan Signal
	MACAddress MACAddress
	DeviceType DeviceType
}

type receiver struct {
}

type transmitter struct {
	TX chan byte
}

func (t transmitter) Write(p []byte) (n int, err error) {
	t.TX = make(chan byte, 1)
	go func() {
		for _, b := range p {
			t.TX <- b
		}
		close(t.TX)
	}()
	n = 0

	err = nil

	return
}

func (d *Device) Start() {
	go run()
}

func run() {
	for {
		return
	}
}

type Power int

type Signal struct {
	B byte
	P Power
}

type MACAddress net.HardwareAddr

type DeviceType int

const (
	FFD DeviceType = iota
	RFD
)
