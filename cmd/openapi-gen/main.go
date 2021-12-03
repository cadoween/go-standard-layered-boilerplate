package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path"

	"github.com/ghodss/yaml"
	
	"github.com/KrisCatDog/go-standard-layered-boilerplate/api/openapi"
)

func main() {
	var output string

	flag.StringVar(&output, "path", "api/openapi", "Path to use for generating OpenAPI 3 files")
	flag.Parse()

	if output == "" {
		log.Fatalln("Target path is required")
	}

	swagger := openapi.NewOpenAPI3()

	// openapi3.json
	data, err := json.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}

	if err := os.WriteFile(path.Join(output, "openapi3.json"), data, 0555); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	// openapi3.yaml
	data, err = yaml.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}

	if err := os.WriteFile(path.Join(output, "openapi3.yaml"), data, 0555); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	log.Println("OpenAPI files successfully generated")
}
