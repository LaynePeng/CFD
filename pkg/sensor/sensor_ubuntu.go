package sensor

import (
	"encoding/json"
	"github.com/LaynePeng/cfd/pkg/spec"
	"github.com/LaynePeng/cfd/pkg/utils"
	"os"
	"strings"
)

type GpuSensorUbuntu struct {
}

func (gs *GpuSensorUbuntu) IsSupported() (bool, error) {
	if _, err := os.Stat("/usr/bin/nvidia-smi"); err == nil {
		return true, nil
	}

	return false, nil
}

func (gs *GpuSensorUbuntu) Desc() (string, error) {
	return "", nil
}

func (gs *GpuSensorUbuntu) Detail() (string, error) {
	ret := utils.RunCmd("nvidia-smi --query-gpu=uuid,gpu_name,memory.free,memory.used,utilization.gpu,driver_version --format=csv,noheader,nounits")
	gpuInfosRet := utils.ParseCSV(ret)

	var gpuInfos []*spec.GPU
	var gpu *spec.GPU
	for _, oneGpuInfo := range gpuInfosRet {
		if len(oneGpuInfo) == 6 {
			gpu = &spec.GPU{
				UUID:          strings.TrimSpace(oneGpuInfo[0]),
				Name:          strings.TrimSpace(oneGpuInfo[1]),
				MemFree:       strings.TrimSpace(oneGpuInfo[2]),
				MemUsed:       strings.TrimSpace(oneGpuInfo[3]),
				GPUUtil:       strings.TrimSpace(oneGpuInfo[4]),
				DriverVersion: strings.TrimSpace(oneGpuInfo[5]),
			}

			gpuInfos = append(gpuInfos, gpu)
		}
	}

	b, _ := json.Marshal(gpuInfos)

	return string(b), nil
}
