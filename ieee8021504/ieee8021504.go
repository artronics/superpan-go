package ieee8021504

import "github.com/artronics/superpan/ieee8021504/mlme"

type IEEE8021504 struct {
	MLME *mlme.MLME
}

func New() *IEEE8021504 {
	i := &IEEE8021504{
		MLME: mlme.New(),
	}

	return i
}
