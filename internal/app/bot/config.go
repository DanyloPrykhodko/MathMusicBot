package bot

type Config struct {
	Token       string `toml:"token"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost dbname=math_music_bot sslmode=disable",
	}
}
