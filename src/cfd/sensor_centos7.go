package cfd

import ()

type GpuSensor struct {
}

type NVRAMSensor struct {
}

type QATSensor struct {
}

type NICBandwidthSensor struct {
}

func (gs *GpuSensor) IsSupported() (bool, error) {
	ret := RunCmd("/usr/sbin/lspci | grep -i nvidia")

	return ret != "", nil
}

func (gs *GpuSensor) Desc() (string, error) {
	return "", nil
}

func (gs *GpuSensor) Detail() (string, error) {
	return "", nil
}

func (ns *NVRAMSensor) IsSupported() (bool, error) {
	ret := RunCmd("lsblk")

	return ParseAndFoundLineByLine("^nvme.*\\s*\\d*:\\d*\\s*\\d*\\s.*", ret), nil
}

func (ns *NVRAMSensor) Desc() (string, error) {
	return "", nil
}

func (ns *NVRAMSensor) Detail() (string, error) {
	return "", nil
}

func (qs *QATSensor) IsSupported() (bool, error) {
	return false, nil
}

func (qs *QATSensor) Desc() (string, error) {
	return "", nil
}

func (qs *QATSensor) Detail() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensor) IsSupported() (bool, error) {
	return true, nil
}

func (gs *NICBandwidthSensor) Desc() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensor) Detail() (string, error) {
	return "", nil
}
