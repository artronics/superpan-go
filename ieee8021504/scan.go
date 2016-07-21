package ieee8021504

type ScanRequest struct {
	ScanType ScanType
}

func (s ScanRequest) request(primitive interface{}, mlme *mlme) error {
	primitive = primitive.(ScanRequest)

	return nil
}

type ScanType int

const (
	ED ScanType = iota
	Active
	Passive
	Orphan
)
