package mlme

import "github.com/artronics/superpan/ieee8021504/pib"

type MLME struct {
	PIB *pib.PIB
}

func New() *MLME {
	m := &MLME{
		PIB: pib.New(),
	}

	return m
}

func (m *MLME) Request(req Requester) (res interface{}, err error) {
	return req.Request(req, m)
}

type Requester interface {
	Request(mlmePrimitive interface{}, mlme *MLME) (res interface{}, err error)
}
