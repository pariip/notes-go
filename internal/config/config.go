package config

type (
	Config struct {
		Auth       Auth
		Database   Database   `yaml:"database"`
		I18n       I18n       `yaml:"i18n"`
		Logger     Logger     `yaml:"logger"`
		User       User       `yaml:"user"`
		Validation Validation `yaml:"validation"`
	}

	Auth struct {
		AccessExpirationInMinute  int    `yaml:"access_expiration_in_minute"`
		RefreshExpirationInMinute int    `yaml:"refresh_expiration_in_minute"`
		JWTSecret                 string `yaml:"jwt_secret"`
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

	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}

	Logger struct {
		MaxAge          string `yaml:"max_age"`
		MaxSize         string `yaml:"max_size"`
		FilenamePattern string `yaml:"filename_pattern"`
		RotationTime    string `yaml:"rotation_time"`
		InternalPath    string `yaml:"internal_path"`
		RequestPath     string `yaml:"request_path"`
	}

	User struct {
	}

	Validation struct {
		UsernameMinLength  int `yaml:"user_name_min_length"`
		UsernameMaxLength  int `yaml:"username_max_length"`
		PasswordMinLetters int `yaml:"password_min_letters"`
	}
)
