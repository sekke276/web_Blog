package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoURI   string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	Port       string `envconfig:"PORT" default:"8080"`
	MongoDB    string `envconfig:"MONGO_DB" default:"test"`
	AuthConfig struct {
		JWTSecret     string `envconfig:"JWT_SECRET" default:"my-secret"`
		JWTExpiration int64  `envconfig:"JWT_EXPIRATION" default:"3600"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
