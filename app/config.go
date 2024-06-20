package app

type Postgres struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type API struct {
	BaseURL string `yaml:"url"`
}

type Config struct {
	Postgres Postgres
	API      API `yaml:"api"`
}
