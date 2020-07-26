package server

type Config struct {
	HttpAddr  string `toml:"http_addr"`
	HttpsAddr string `toml:"https_addr"`
	LogLebel  string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		HttpAddr:  ":8080",
		HttpsAddr: ":8443",
		LogLebel:  "debug",
	}
}
