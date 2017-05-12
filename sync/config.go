package sync

type Config struct {
	MySQL string
}

func NewConfig() *Config {
	cfg := &Config{
		MySQL: "",
	}

	return cfg
}
