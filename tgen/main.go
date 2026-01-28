package main

import (
	"fmt"
	"os"
	"tgen/types"

	"github.com/goccy/go-yaml"
)

func main() {
	f, err := os.ReadFile("./models.yml")
	if err != nil {
		panic(err)
	}

	var manifest types.Manifest
	if err := yaml.Unmarshal(f, &manifest); err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", manifest.Config)
	fmt.Printf("%+v\n", manifest.Models)

	// config := manifest.Config

	// models := manifest.Models.ToArray(config.Defaults)
	// fmt.Printf("%+v", models)
}
