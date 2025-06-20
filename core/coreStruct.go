package core

type AppConfig struct {
	Account AccountConfig `yaml:"config"`
}

type AccountConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
