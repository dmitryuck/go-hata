package config

var (
	Instance *Config
)

// Config declaration
type Config struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	MongoDBUrl string `toml:"mongodb_url"`
	DBName     string `toml:"db_name"`
}

// New creates default config object
func New() *Config {
	Instance = &Config{
		BindAddr:   ":80",
		LogLevel:   "debug",
		MongoDBUrl: "mongodb://localhost:27017",
		DBName:     "hatadb",
	}

	return Instance
}
