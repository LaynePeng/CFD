package spec

import (
	"encoding/json"
	"fmt"
	"log"
)

type NIC struct {
	Device    string `json:"device"`
	Bandwidth int64  `json:"bandwidth"`
}

type GPU struct {
	UUID          string `json:"uuid"`
	Name          string `json:"name"`
	MemFree       string `json:"mem_free(MiB)"`
	MemUsed       string `json:"mem_used(MiB)"`
	GPUUtil       string `json:"gpu_util(%)"`
	DriverVersion string `json:"driver_version`
}

type NVRAM struct {
	Name       string `json:"type"`
	Major      string `json:"maj"`
	Min        string `json:"min"`
	Size       string `json:"size"`
	MountPoint string `json:"MountPoint"`
}

type QAT struct {
	Device string `json:"device"`
	Type   string `json:"type"`
	NodeId string `json:"node"`
	Engine string `json:"engine"`
	State  string `json:"state"`
}

func (n *NIC) ToJson() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func (n *NIC) FromJson(nicJson string) *NIC {
	if err := json.Unmarshal([]byte(nicJson), n); err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	return n
}

func (g *GPU) ToJson() string {
	b, _ := json.Marshal(g)
	return string(b)
}

func (g *GPU) FromJson(gpuJson string) *GPU {
	if err := json.Unmarshal([]byte(gpuJson), g); err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	return g
}

func (nm *NVRAM) ToJson() string {
	b, _ := json.Marshal(nm)
	return string(b)
}

func (nm *NVRAM) FromJson(nvramJson string) *NVRAM {
	if err := json.Unmarshal([]byte(nvramJson), nm); err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	return nm
}

func (q *QAT) ToJson() string {
	b, _ := json.Marshal(q)
	return string(b)
}

func (q *QAT) FromJson(qatJson string) *QAT {
	if err := json.Unmarshal([]byte(qatJson), q); err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	return q
}
