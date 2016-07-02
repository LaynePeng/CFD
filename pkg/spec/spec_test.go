package spec

import (
	"testing"
)

func TestNicToJson(t *testing.T) {
	nic := &NIC{
		Device:    "eth0",
		Bandwidth: 1000,
	}

	t.Log(nic.ToJson())
}
