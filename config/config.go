package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
	Dialect  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "localhost",
			Port:     "5433",
			Username: "postgres",
			Password: "changepassword",
			Name:     "contact_book",
			Charset:  "utf8",
			Dialect:  "postgres",
		},
	}
}
