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
	App      AppConfig
	Postgres Postgres
	API      API `yaml:"api"`
}

type AppConfig struct {
	Hosts Hosts `yaml:"hosts"`
}

type Hosts struct {
	Images string `yaml:"images"`
}
