package config

type Config struct {
	Server struct {
		Port     string `yaml:"port" envconfig:"SERVER_PORT"`
		BasePath string `yaml:"basePath" envconfig:"SERVER_BASEPATH"`
	} `yaml:"server"`
	Database struct {
		Url     string `yaml:"url" envconfig:"DATABASE_URL"`
		MaxOpen int    `yaml:"maxOpen" envconfig:"DATABASE_MAXOPEN"`
		MaxIdle int    `yaml:"maxIdle" envconfig:"DATABASE_MAXIDLE"`
	} `yaml:"database"`
}
