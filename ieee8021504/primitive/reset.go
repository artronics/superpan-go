package primitive

import "github.com/artronics/superpan/ieee8021504/mlme"

type ResetRequest struct {
	SetDefaultPIB bool
}

func (s ResetRequest) request(primitive interface{}, mlme *mlme.MLME) error {
	primitive = primitive.(ResetRequest)

	return nil
}
