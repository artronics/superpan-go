package node

type Device struct {
	DeviceType DeviceType
}

type DeviceType int

const (
	FFD DeviceType = iota
	RFD
)
