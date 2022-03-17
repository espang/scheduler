package system

type Config struct{}

func ConfigFromEnv() *Config {
	return &Config{}
}
