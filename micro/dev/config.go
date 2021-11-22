package main

type ConfigEnv struct {
	Runs        []string `json:"runs"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	MicroAuth   string   `json:"micro_auth"`
	MicroSecret string   `json:"micro_secret"`
	MicroProxy  string   `json:"micro_proxy"`
	ConfigFile  string   `json:"config_file"`
}
