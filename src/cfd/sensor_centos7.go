package cfd

import (
	"fmt"
	"strconv"
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

	return ret, nil
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

	return ReturnAndFoundLineByLine("^nvme.*\\s*\\d*:\\d*\\s*\\d*\\s.*", ret), nil
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
