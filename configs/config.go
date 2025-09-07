package configs

type Config struct {
	Smtp SmtpConfig
}

type SmtpConfig struct {
	Email    string
	Password string
	Address  string
}

func NewConfig() *Config {
	return &Config{
		Smtp: SmtpConfig{
			Email:    "pub.nes@mail.ru",
			Password: "123",
			Address:  "123",
		},
	}
}
