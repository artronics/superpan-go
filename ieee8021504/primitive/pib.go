package primitive

import (
	"errors"
	"github.com/artronics/superpan/ieee8021504/mlme"
	"github.com/artronics/superpan/ieee8021504/pib"
)

type GetRequest struct {
	PIBAttribute pib.PIBAttribute
}
type GetResponse struct {
	Status            Status
	PIBAttribute      pib.PIBAttribute
	PIBAttributeValue interface{}
}

func (s GetRequest) Request(primitive interface{}, mlme *mlme.MLME) (interface{}, error) {
	primitive = primitive.(GetRequest)

	r := GetResponse{PIBAttribute: s.PIBAttribute}
	a, ok := mlme.PIB.Attribute[s.PIBAttribute]

	if !ok {
		r.Status = UNSUPPORTED_ATTRIBUTE
		r.PIBAttributeValue = nil
		return r, errors.New("UNSUPPORTED_ATTRIBUTE")
	}
	r.Status = SUCCESS
	r.PIBAttributeValue = a

	return r, nil
}

type Status int

const (
	SUCCESS Status = iota
	UNSUPPORTED_ATTRIBUTE
)
