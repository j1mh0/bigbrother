package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Project string
	Author  string `yaml:"author"`
	Version float32
	Content []string
}

func main() {

	configFile, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal("read config file error:", err)
		return
	}

	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("read config struct error:", err)
		return
	}

	fmt.Printf("Project: %s\nAuthor: %s\nVersion: %g\n", config.Project, config.Author, config.Version)

	for _, value := range config.Content {
		fmt.Println(value)
	}

}
