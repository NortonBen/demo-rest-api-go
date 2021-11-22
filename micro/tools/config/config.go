package main

import (
	"apm/pkg/util"
	"fmt"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := cli.App{
		Name: "config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				EnvVars: []string{"FILE_CONFIG"},
				Value:   "config.json",
			},
		},
		Before: func(context *cli.Context) error {
			srv := service.New(
				service.Name("config-service"),
				service.Version("0.0.1"),
			)

			srv.Init()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "store",
				Action: func(c *cli.Context) error {
					fileName := c.String("file")

					data := make(map[string]interface{})
					err := util.ReadFile(&data, fileName)
					if err != nil {
						return nil
					}

					writeData(data, "")
					return nil
				},
			},
			{
				Name: "get",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					val, err := config.Get(name)
					if err != nil {
						logger.Error(err)
					}
					logger.Info(val)

					return nil
				},
			},
			{
				Name: "set",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					value := c.Args().Get(1)

					err := config.Set(name, value)
					if err != nil {
						logger.Error(err)
					}
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

func writeData(data map[string]interface{}, root string) {
	for key, val := range data {
		keyNew := key
		if root != "" {
			keyNew = fmt.Sprintf("%s.%s", root, key)
		}

		switch val.(type) {
		case map[string]interface{}:
			writeData(val.(map[string]interface{}), keyNew)
		default:
			config.Set(keyNew, val)
		}
	}

}
