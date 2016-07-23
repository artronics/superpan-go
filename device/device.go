package device

import (
	"github.com/artronics/superpan/medium"
	"net"
)

type Device struct {
	receiver
	transmitter
	medium     *medium.Medium
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

type MACAddress net.HardwareAddr

type DeviceType int

const (
	FFD DeviceType = iota
	RFD
)
