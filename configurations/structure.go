package configurations

import "time"

type Structure struct {
	Service struct {
		Name     string `yaml:"name"`
		Debug    string `yaml:"debug"`
		TimeZone string `yaml:"timezone"`
	} `yaml:"service"`
	Server struct {
		Port            string        `yaml:"port"`
		Host            string        `yaml:"host"`
		ReadTimeout     time.Duration `yaml:"write_timeout"`
		WriteTimeout    time.Duration `yaml:"read_timeout"`
		ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	} `yaml:"server"`
	Database struct {
		Port       string        `yaml:"port"`
		Host       string        `yaml:"host"`
		Name       string        `yaml:"name"`
		Collection string        `yaml:"collection"`
		Timeout    time.Duration `yaml:"timeout"`
	} `yaml:"database"`
}
