package config

var (
	Instance *Config
)

// Config declaration
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig creates default config object
func NewConfig() *Config {
	Instance = &Config{
		BindAddr: ":80",
		LogLevel: "debug",
	}

	return Instance
}
