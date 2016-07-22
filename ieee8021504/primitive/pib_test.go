package primitive

import (
	"github.com/artronics/superpan/ieee8021504"
	"github.com/artronics/superpan/ieee8021504/pib"
	"testing"
)

func TestPib(t *testing.T) {
	i := ieee8021504.New()
	r, err := i.MLME.Request(GetRequest{PIBAttribute: pib.MacExtendedAddress})
	v := r.(GetResponse)
	//success scenario
	if v.Status != 0 {
		t.Error("expected 0 (SUCCESS) but got %v", v.Status)
	}
	if err != nil {
		t.Error("expected no error")
	}

	//Fail scenario
	r, err = i.MLME.Request(GetRequest{PIBAttribute: "foo"})
	v = r.(GetResponse)
	if v.Status != 1 {
		t.Error("expected 1 (UNSUPPORTED_ATTRIBUTE) but got %v", v.Status)
	}
	if err == nil {
		t.Error("expected error")
	}
}
