package primitive

import (
	"github.com/artronics/superpan/ieee8021504/mlme"
)

type ScanRequest struct {
	ScanType ScanType
}
type ScanResponse struct {
}

func (s ScanRequest) Request(primitive interface{}, mlme *mlme.MLME) (res interface{}, err error) {
	primitive = primitive.(ScanRequest)
	res = ScanResponse{}
	err = nil
	return
}

type ScanType int

const (
	ED ScanType = iota
	Active
	Passive
	Orphan
)
