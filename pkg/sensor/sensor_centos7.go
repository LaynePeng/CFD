package sensor

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/LaynePeng/cfd/pkg/spec"
	"github.com/LaynePeng/cfd/pkg/utils"
)

type GpuSensorCentOS7 struct {
}

type NVRAMSensorCentOS7 struct {
}

type QATSensorCentOS7 struct {
}

type FPGASensorLinux struct {
}

type NICBandwidthSensorCentOS7 struct {
}

func (gs *GpuSensorCentOS7) IsSupported() (bool, error) {
	ret := utils.RunCmd("/usr/sbin/lspci | grep -i nvidia")

	return ret != "", nil
}

func (gs *GpuSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (gs *GpuSensorCentOS7) Detail() (string, error) {
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

func (ns *NVRAMSensorCentOS7) IsSupported() (bool, error) {
	ret := utils.RunCmd("lsblk")

	return utils.ParseAndFoundLineByLine("^nvme.*\\s*\\d*:\\d*\\s*\\d*\\s.*", ret), nil
}

func (ns *NVRAMSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (ns *NVRAMSensorCentOS7) Detail() (string, error) {
	ret := utils.RunCmd("lsblk")
	regexp_for_nvram := "^(nvme\\S*)\\s+(\\d*):(\\d*)\\s+\\d+\\s+(\\d+\\w{1})\\s+\\d+\\s+\\w+(.*)"
	rvramInfosRet := utils.ReturnSubValueOfFoundLineByLine(regexp_for_nvram, ret)

	var nvramInfos []*spec.NVRAM
	var nvram *spec.NVRAM
	for _, oneRvramInfos := range rvramInfosRet {
		if len(oneRvramInfos) > 5 {
			nvram = &spec.NVRAM{
				Name:       strings.TrimSpace(oneRvramInfos[1]),
				Major:      strings.TrimSpace(oneRvramInfos[2]),
				Min:        strings.TrimSpace(oneRvramInfos[3]),
				Size:       strings.TrimSpace(oneRvramInfos[4]),
				MountPoint: strings.TrimSpace(oneRvramInfos[5]),
			}

			nvramInfos = append(nvramInfos, nvram)
		}
	}

	b, _ := json.Marshal(nvramInfos)

	return string(b), nil
}

func (qs *QATSensorCentOS7) IsSupported() (bool, error) {
	ret := utils.RunCmd("service qat_service status")

	return utils.ParseAndFoundLineByLine(".+, state=up", ret), nil
}

func (qs *QATSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (qs *QATSensorCentOS7) Detail() (string, error) {
	ret := utils.RunCmd("service qat_service status")
	regexp_for_qat := "^(.+)\\s+-\\s+type=(\\S+),\\s+inst_id=0,\\s+node_id=(\\d+),\\s+bdf=.+,\\s+#accel=\\d+,\\s+#engines=(\\d+),\\s+state=(\\w+)"
	qatInfosRet := utils.ReturnSubValueOfFoundLineByLine(regexp_for_qat, ret)

	var qatInfos []*spec.QAT
	var qat *spec.QAT
	for _, oneQatInfos := range qatInfosRet {
		if len(oneQatInfos) > 5 {
			qat = &spec.QAT{
				Device: strings.TrimSpace(oneQatInfos[1]),
				Type:   strings.TrimSpace(oneQatInfos[2]),
				NodeId: strings.TrimSpace(oneQatInfos[3]),
				Engine: strings.TrimSpace(oneQatInfos[4]),
				State:  strings.TrimSpace(oneQatInfos[5]),
			}

			qatInfos = append(qatInfos, qat)
		}
	}

	b, _ := json.Marshal(qatInfos)

	return string(b), nil
}

func (fs *FPGASensorLinux) IsSupported() (bool, error) {
	ret := utils.RunCmd("/usr/bin/xbsak list")
	lines := strings.Split(ret, "\n")

	return len(lines) > 2, nil
}

func (fs *FPGASensorLinux) Desc() (string, error) {
	return "", nil
}

func (fs *FPGASensorLinux) Detail() (string, error) {
	ret := utils.RunCmd("/usr/bin/xbsak list")
	lines := strings.Split(ret, "\n")

	var fpgaInfos []*spec.FPGA
	var fpga *spec.FPGA
	id := 0
	for i := 1; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			tunples := strings.Split(lines[i], " ")
			if len(tunples) == 2 {
				fpga = &spec.FPGA{
					IP:       strings.TrimSpace("ANY"),
					Type:     strings.TrimSpace("FPGA"),
					ID:       uint64(id),
					Platform: strings.TrimSpace("Xilinx"),
					Device:   strings.TrimSpace(tunples[1]),
				}
				id++
			}

			fpgaInfos = append(fpgaInfos, fpga)
		}
	}

	b, _ := json.Marshal(fpgaInfos)

	return string(b), nil
}

func (gs *NICBandwidthSensorCentOS7) IsSupported() (bool, error) {
	return true, nil
}

func (gs *NICBandwidthSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensorCentOS7) Detail() (string, error) {
	var speeds []int

	ret := utils.RunCmd("ifconfig")
	regexp_for_nic := "^(en.*|eth\\d+):\\s+flags=\\d.*"
	regexp_for_speed := "\\s*Speed:\\s+(\\d+).*/s"

	nicInfo := utils.ReturnSubValueOfFoundLineByLine(regexp_for_nic, ret)

	for _, oneNicInfo := range nicInfo {
		nic_name := oneNicInfo[1]
		cmd_for_nic_speed := fmt.Sprintf("%s %s", "ethtool", nic_name)

		ret = utils.RunCmd(cmd_for_nic_speed)
		speed_values := utils.ReturnSubValueOfFoundLineByLine(regexp_for_speed, ret)
		if speed_values != nil && len(speed_values) > 0 {
			nic_speed := speed_values[0][1]
			i, err := strconv.Atoi(nic_speed)
			if err == nil {
				speeds = append(speeds, i)
			}
		}
	}

	maxSpeed := 0
	for _, theSpeed := range speeds {
		if maxSpeed < theSpeed {
			maxSpeed = theSpeed
		}
	}

	return strconv.Itoa(maxSpeed), nil
}
