package sensor

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var (
	CENTOS_7             = "centos7"
	UBUNTU               = "ubuntu"
	HardwareFunctionsSet = []string{"gpu", "nvram", "qat"}
)

type Sensor interface {
	IsSupported() (bool, error)
	Desc() (string, error)
	Detail() (string, error)
}

func NewGpuSensor() Sensor {
	switch checkOSVersion() {
	case CENTOS_7:
		return &GpuSensorCentOS7{}
	case UBUNTU:
		return &GpuSensorUbuntu{}
	}
	return &GpuSensorFake{}
}

func NewNVRAMSensor() Sensor {
	switch checkOSVersion() {
	case CENTOS_7:
		return &NVRAMSensorCentOS7{}
	}
	return &NVRAMSensorFake{}
}

func NewQATSensor() Sensor {
	switch checkOSVersion() {
	case CENTOS_7:
		return &QATSensorCentOS7{}
	}
	return &QATSensorFake{}
}

func NewNICBandwidthSensor() Sensor {
	switch checkOSVersion() {
	case CENTOS_7:
		return &NICBandwidthSensorCentOS7{}
	}

	return &NICBandwidthSensorFake{}
}

func checkOSVersion() string {
	if _, err := os.Stat("/etc/redhat-release"); err == nil {
		f, err := os.Open("/etc/redhat-release")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			if strings.Contains(line, "CentOS Linux release 7") {
				return CENTOS_7
			}
		}
	} else if _, err := os.Stat("/etc/issue"); err == nil {
		f, err := os.Open("/etc/issue")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			if strings.Contains(line, "Ubuntu") {
				return UBUNTU
			}
		}
	} else {
		log.Println("Your OS is not supported. Use FAKE Driver")
	}

	return "fake"
}
