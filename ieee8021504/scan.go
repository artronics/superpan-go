package ieee8021504

import ()
import "fmt"

type ScanRequest struct {
	ScanType ScanType
}

func (s ScanRequest) request(primitive interface{}, mlme *mlme) error {
	primitive = primitive.(ScanRequest)
	if mlme.DeviceType == RFD && (s.ScanType == ED) {
		panic("ED scan is not supported in RFD")
	}

	return nil
}

type ScanType int

const (
	ED ScanType = iota
	Active
	Passive
	Orphan
)
