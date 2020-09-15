package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

//Config type
type Config struct {
	R2Commands []struct {
		Cmd  string `yaml:"cmd"`
		Idx  string `yaml:"idx"`
		File *os.File
	} `yaml:"r2commands"`
}

func loadConfigFromFile(filePath string, config *Config) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed opening file: %s, %v\n", filePath, err)
		return err
	}
	defer func() { _ = configFile.Close() }()
	yamlParser := yaml.NewDecoder(configFile)
	err = yamlParser.Decode(&config)
	if err == io.EOF {
		err = nil
	}
	return err
}

func loadConfig(configFile string) error {

	loadConfigFromFile(configFile, &config)
	return nil
}
