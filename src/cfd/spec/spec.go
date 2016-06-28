package spec

import (
	"encoding/json"
)

type NIC struct {
	Device    string `json:"device"`
	Bandwidth int64  `json:"bandwidth"`
}

type GPU struct {
	Type string `json:"type"`
	Desc string `json:"desc"`
}

type NVRAM struct {
	Name       string `json:"type"`
	Major      string `json:"maj"`
	Min        string `json:"min"`
	Size       string `json:"size"`
	MountPoint string `json:"MountPoint"`
}

type QAT struct {
}

func (n *NIC) ToJson() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func (g *GPU) ToJson() string {
	b, _ := json.Marshal(g)
	return string(b)
}

func (nm *NVRAM) ToJson() string {
	b, _ := json.Marshal(nm)
	return string(b)
}

func (q *QAT) ToJson() string {
	b, _ := json.Marshal(q)
	return string(b)
}
