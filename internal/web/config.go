package web

type Config struct {
	HttpAddr  string `toml:"web_http_addr"`
	HttpsAddr string `toml:"web_https_addr"`
	LogLebel  string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		HttpAddr:  ":80",
		HttpsAddr: ":443",
		LogLebel:  "debug",
	}
}
