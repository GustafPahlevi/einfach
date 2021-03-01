package configuration

import (
	"log"
	"os"

	"github.com/GustafPahlevi/go-simple-svc/configurations"

	"gopkg.in/yaml.v2"
)

const (
	defaultPath = "./configurations/service.yaml"
	//defaultPath = "../app/configurations/service.yaml"
)

// Read is a kit that will read the configuration
func Read(path string) (config configurations.Structure, err error) {
	if path == "" {
		path = defaultPath
	}

	file, err := os.Open(path)
	if err != nil {
		return configurations.Structure{}, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("error while close file, got: %v", err)
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return configurations.Structure{}, err
	}

	return config, nil
}
