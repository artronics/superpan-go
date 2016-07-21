package ieee8021504

import (
	"testing"
)

func TestScan_request_should_panic_if_type_assertion_fails(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but got none.")
		}
	}()
	s := ScanRequest{}
	s.request(requesterMock{}, nil)
}

func TestScan_validate_ED_is_not_supported_in_RFD(t *testing.T) {
	defer shouldPanic(t)
	i := IEEE8021504{}
	i.MLME.Request(ScanRequest{ScanType: ED})
}
func shouldPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Error("Expected panic but got none.")
	}
}
