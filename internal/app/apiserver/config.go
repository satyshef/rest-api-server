package apiserver

type Config struct {
	Addr         string `toml:"bind_addr"`
	ReadTimeout  int    `toml:"read_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
	LogLevel     string `toml:"log_level"`
}

// Create config instance
func MakeConfig() *Config {
	return &Config{
		Addr:         ":8080",
		ReadTimeout:  10,
		WriteTimeout: 10,
		LogLevel:     "debug",
	}
}
