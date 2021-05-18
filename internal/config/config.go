package config

type (
	Config struct {
		Database Database `yaml:"database"`
	}
	Database struct {
		Postgres Postgres `yaml:"postgres"`
	}

	Postgres struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		DBName    string `yaml:"db_name"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		SSLMode   string `yaml:"ssl_mode"`
		TimeZone  string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
		Migration bool   `yaml:"migration"`
	}
)
