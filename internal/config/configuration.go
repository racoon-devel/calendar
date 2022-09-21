package config

type Database struct {
	Name     string `default:"calendar" env:"DB_NAME"`
	User     string `default:"calendar" env:"DB_USER"`
	Password string `required:"true" env:"DB_PASSWORD"`
	Host     string `default:"127.0.0.1" env:"DB_HOST"`
	Port     uint   `default:"5432" env:"DB_PORT"`
}

type Configuration struct {
	Database Database
	Http     struct {
		Endpoint string `default:"0.0.0.0:8080" env:"HTTP_ENDPOINT"`
	}
}
