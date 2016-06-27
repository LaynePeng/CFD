package cfd

import ()

type GpuSensorFake struct {
}

type NVRAMSensorFake struct {
}

type QATSensorFake struct {
}

type NICBandwidthSensorFake struct {
}

func (gs *GpuSensorFake) IsSupported() (bool, error) {
	return true, nil
}

func (gs *GpuSensorFake) Desc() (string, error) {
	return "", nil
}

func (gs *GpuSensorFake) Detail() (string, error) {
	return "", nil
}

func (ns *NVRAMSensorFake) IsSupported() (bool, error) {
	return true, nil
}

func (ns *NVRAMSensorFake) Desc() (string, error) {
	return "", nil
}

func (ns *NVRAMSensorFake) Detail() (string, error) {
	return "", nil
}

func (qs *QATSensorFake) IsSupported() (bool, error) {
	return false, nil
}

func (qs *QATSensorFake) Desc() (string, error) {
	return "", nil
}

func (qs *QATSensorFake) Detail() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensorFake) IsSupported() (bool, error) {
	return true, nil
}

func (gs *NICBandwidthSensorFake) Desc() (string, error) {
	return "", nil
}

func (gs *NICBandwidthSensorFake) Detail() (string, error) {
	return "", nil
}
