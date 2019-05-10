package main

const ConfigPath = "./local.json"

func main() {
	config, err := LoadConfigFromFileOrDefault(ConfigPath)
	if err != nil {
		panic(err)
	}

	server := NewServer(config)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
