package cfd

import (
	"cfd/spec"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type GpuSensorCentOS7 struct {
}

type NVRAMSensorCentOS7 struct {
}

type QATSensorCentOS7 struct {
}

type NICBandwidthSensorCentOS7 struct {
}

func (gs *GpuSensorCentOS7) IsSupported() (bool, error) {
	ret := RunCmd("/usr/sbin/lspci | grep -i nvidia")

	return ret != "", nil
}

func (gs *GpuSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (gs *GpuSensorCentOS7) Detail() (string, error) {
	ret := RunCmd("/usr/sbin/lspci | grep -i nvidia")
	regexp_for_parse_gpu := ".*3D controller:\\s*(.*)\\[(.*)\\]"
	gpuInfosRet := ReturnSubValueOfFoundLineByLine(regexp_for_parse_gpu, ret)

	var gpuInfos []*spec.GPU
	var gpu *spec.GPU
	for _, oneGpuInfo := range gpuInfosRet {
		if len(oneGpuInfo) == 3 {
			gpu = &spec.GPU{
				Type: strings.TrimSpace(oneGpuInfo[2]),
				Desc: strings.TrimSpace(oneGpuInfo[1]),
			}

			gpuInfos = append(gpuInfos, gpu)
		}
	}

	b, _ := json.Marshal(gpuInfos)

	return string(b), nil
}

func (ns *NVRAMSensorCentOS7) IsSupported() (bool, error) {
	ret := RunCmd("lsblk")

	return ParseAndFoundLineByLine("^nvme.*\\s*\\d*:\\d*\\s*\\d*\\s.*", ret), nil
}

func (ns *NVRAMSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (ns *NVRAMSensorCentOS7) Detail() (string, error) {
	ret := RunCmd("lsblk")
	regexp_for_nvram := "^(nvme\\S*)\\s+(\\d*):(\\d*)\\s+\\d+\\s+(\\d+\\w{1})\\s+\\d+\\s+\\w+(.*)"
	rvramInfosRet := ReturnSubValueOfFoundLineByLine(regexp_for_nvram, ret)

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
	return false, nil
}

func (qs *QATSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (qs *QATSensorCentOS7) Detail() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensorCentOS7) IsSupported() (bool, error) {
	return true, nil
}

func (gs *NICBandwidthSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensorCentOS7) Detail() (string, error) {
	var speeds []int

	ret := RunCmd("ifconfig")
	regexp_for_nic := "(en.*|eth\\d+):\\s+flags=\\d.*"
	regexp_for_speed := "\\s*Speed:\\s+(\\d+).*/s"

	nicInfo := ReturnSubValueOfFoundLineByLine(regexp_for_nic, ret)

	for _, oneNicInfo := range nicInfo {
		nic_name := oneNicInfo[1]
		cmd_for_nic_speed := fmt.Sprintf("%s %s", "ethtool", nic_name)

		ret = RunCmd(cmd_for_nic_speed)
		speed_values := ReturnSubValueOfFoundLineByLine(regexp_for_speed, ret)
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