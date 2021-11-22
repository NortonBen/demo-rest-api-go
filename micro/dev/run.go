package main

import (
	"apm/micro/run"
	"apm/pkg/util"
	"github.com/micro/micro/v3/service/config"
	"log"
	"os"
)

var configEnv ConfigEnv

func init() {

	var fileEnv = ".dev.env.json"

	env, success := os.LookupEnv("ENV_APM")
	if success {
		fileEnv = env + fileEnv
	}

	err := util.ReadFile(&configEnv, fileEnv)
	if err != nil {
		log.Fatalln(env)
	}

	os.Setenv("MICRO_SERVICE_NAME", configEnv.Name)
	os.Setenv("MICRO_SERVICE_ADDRESS", configEnv.Address)
	os.Setenv("MICRO_AUTH_ID", configEnv.MicroAuth)
	os.Setenv("MICRO_AUTH_SECRET", configEnv.MicroSecret)
	os.Setenv("MICRO_PROXY", configEnv.MicroProxy)

	config.DefaultConfig, err = NewConfig(configEnv.ConfigFile)
	if err != nil {
		log.Fatalln(env)
	}
}

func main() {
	run.Run()
}
