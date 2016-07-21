package ieee8021504

import "testing"

func TestMlme_Request(t *testing.T) {
}

type requesterMock struct {
}

func (r requesterMock) request(primitive interface{}) error {
	return nil
}
