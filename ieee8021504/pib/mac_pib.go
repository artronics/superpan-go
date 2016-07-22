package pib

import "net"

type PIBAttribute string

type PIB struct {
	Attribute map[PIBAttribute]interface{}
}

func New() *PIB {
	p := &PIB{Attribute: make(map[PIBAttribute]interface{})}
	p.Attribute[MacExtendedAddress] = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	return p
}

const (
	MacExtendedAddress PIBAttribute = "macExtendedAddress"
)

type macExtendedAddress net.HardwareAddr
