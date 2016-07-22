package primitive

import (
	"github.com/artronics/superpan/ieee8021504"
	"testing"
)

func TestScan_request_should_panic_if_type_assertion_fails(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but got none.")
		}
	}()
	s := ScanRequest{}
	s.Request(requesterMock{}, nil)
}

func TestScan_validate_ED_is_not_supported_in_RFD(t *testing.T) {
	defer shouldPanic(t)
	i := ieee8021504.IEEE8021504{}
	i.MLME.Request(ScanRequest{ScanType: ED})
}
func shouldPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Error("Expected panic but got none.")
	}
}

type requesterMock struct {
}

func (r requesterMock) Request(primitive interface{}) (interface{}, error) {
	return nil, nil
}
