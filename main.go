package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"cfd"
	"github.com/urfave/cli"
)

func listAllSupports() ([]string, error) {
	supports := []string{}
	for _, f := range cfd.HardwareFunctionsSet {
		supported, err := testFunctionSupported(f)
		if err != nil {
			return nil, err
		}
		if supported {
			supports = append(supports, f)
		}
	}
	return supports, nil
}

func showFunctionDetail() (string, error) {
	return "", nil
}

func testFunctionSupported(name string) (bool, error) {
	var sensor cfd.Sensor
	switch name {
	case "gpu":
		sensor = cfd.NewGpuSensor()
	case "nvram":
		sensor = cfd.NewNVRAMSensor()
	case "qat":
		sensor = cfd.NewQATSensor()
	case "nic_bandwidth":
		sensor = cfd.NewNICBandwidthSensor()
	default:
		return false, errors.New("Not a valid function!")
	}

	return sensor.IsSupported()
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = "v1.0.0"
	app.Name = "Composable Function Discover"
	app.Usage = "A toolkit for discovering hardware functons in system"

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List the harware functions supported",
			Action: func(c *cli.Context) error {
				supports, _ := listAllSupports()
				fmt.Println(strings.Join(supports, ","))
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show the detal of one hardware function",
			Action: func(c *cli.Context) error {
				fmt.Println("Show")
				return nil
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Test if a hardware function supported",
			Action: func(c *cli.Context) error {
				isSupport, _ := testFunctionSupported(c.String("function"))
				fmt.Println(isSupport)
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "function, f",
					Usage: "Specify a hardware function to test",
				},
			},
		},
	}

	app.Run(os.Args)
}
