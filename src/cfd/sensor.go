package cfd

import ()

var (
	HardwareFunctionsSet = []string{"gpu", "nvram", "qat"}
)

type Sensor interface {
	IsSupported() (bool, error)
	Desc() (string, error)
	Detail() (string, error)
}
