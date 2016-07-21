package main

import (
	"github.com/artronics/superpan/ieee8021504"
)

func main() {
	i := ieee8021504.IEEE8021504{}
	i.MLME.Request(ieee8021504.ScanRequest{})
}
