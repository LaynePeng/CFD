package spec

import (
	"testing"
)

func TestNicJson(t *testing.T) {
	nic := &NIC{
		Device:    "eth0",
		Bandwidth: 1000,
	}

	nicJson := nic.ToJson()

	t.Log(nicJson)

	nic2 := &NIC{}
	nic2 = nic2.FromJson(nicJson)

	nic2Json := nic2.ToJson()

	if nicJson != nic2Json {
		t.Fail()
	}
}

func TestGPUJson(t *testing.T) {
	gpu := &GPU{
		Type: "ABC",
		Desc: "Nvidia GPU",
	}

	gpuJson := gpu.ToJson()

	t.Log(gpuJson)

	gpu2 := &GPU{}
	gpu2 = gpu2.FromJson(gpuJson)

	gpu2Json := gpu2.ToJson()

	if gpuJson != gpu2Json {
		t.Fail()
	}
}

func TestNVRAMJson(t *testing.T) {
	nvram := &NVRAM{
		Name:       "NVRAM",
		Major:      "200",
		Min:        "1",
		Size:       "3G",
		MountPoint: "",
	}

	nvramJson := nvram.ToJson()

	t.Log(nvramJson)

	nvram2 := &NVRAM{}
	nvram2 = nvram2.FromJson(nvramJson)

	nvram2Json := nvram2.ToJson()

	if nvramJson != nvram2Json {
		t.Fail()
	}
}

func TestQATJson(t *testing.T) {
	qat := &QAT{
		Device: "QAT Device",
		Type:   "Intel QAT",
		NodeId: "1",
		Engine: "Intel QAT",
		State:  "Up",
	}

	qatJson := qat.ToJson()

	t.Log(qatJson)

	qat2 := &QAT{}
	qat2 = qat2.FromJson(qatJson)

	qat2Json := qat2.ToJson()

	if qatJson != qat2Json {
		t.Fail()
	}
}
