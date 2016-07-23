package device

import "testing"

func TestTransmitter_Write(t *testing.T) {
	d := Device{}
	d.Write([]byte{1, 2})
	b := <-d.TX
	if b != 2 {
		t.Errorf("expected 2 items but got %d", b)
	}
}
