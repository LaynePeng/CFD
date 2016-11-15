package sensor

import (
	"log"
)

type GpuSensorFake struct {
}

type NVRAMSensorFake struct {
}

type QATSensorFake struct {
}

type NICBandwidthSensorFake struct {
}

func (gs *GpuSensorFake) IsSupported() (bool, error) {
	log.Println("Using FAKE Driver")
	return true, nil
}

func (gs *GpuSensorFake) Desc() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (gs *GpuSensorFake) Detail() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (ns *NVRAMSensorFake) IsSupported() (bool, error) {
	log.Println("Using FAKE Driver")
	return true, nil
}

func (ns *NVRAMSensorFake) Desc() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (ns *NVRAMSensorFake) Detail() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (qs *QATSensorFake) IsSupported() (bool, error) {
	log.Println("Using FAKE Driver")
	return false, nil
}

func (qs *QATSensorFake) Desc() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (qs *QATSensorFake) Detail() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (gs *NICBandwidthSensorFake) IsSupported() (bool, error) {
	log.Println("Using FAKE Driver")
	return true, nil
}

func (gs *NICBandwidthSensorFake) Desc() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}

func (gs *NICBandwidthSensorFake) Detail() (string, error) {
	log.Println("Using FAKE Driver")
	return "", nil
}
