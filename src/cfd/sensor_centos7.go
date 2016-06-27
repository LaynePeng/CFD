package cfd

import ()

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
	return "", nil
}

func (ns *NVRAMSensorCentOS7) IsSupported() (bool, error) {
	ret := RunCmd("lsblk")

	return ParseAndFoundLineByLine("^nvme.*\\s*\\d*:\\d*\\s*\\d*\\s.*", ret), nil
}

func (ns *NVRAMSensorCentOS7) Desc() (string, error) {
	return "", nil
}

func (ns *NVRAMSensorCentOS7) Detail() (string, error) {
	return "", nil
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
	return "", nil
}
