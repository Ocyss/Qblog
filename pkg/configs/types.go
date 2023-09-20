package configs

type Conf struct {
	App      App      `yaml:"App"`
	Registry Registry `yaml:"Registry"`
	Db       Db       `yaml:"Db"`
	Oss      Oss      `yaml:"Oss"`
	Push     Push     `yaml:"Push"`
}
type Oss []map[string]string

type App struct {
	Mode string `yaml:"Mode"`
	Jwt  string `yaml:"Jwt"`
}

type Db struct {
	Pgsql Pgsql `yaml:"Pgsql"`
	Redis Redis `yaml:"Redis"`
}

type Registry struct {
	Address  []string `yaml:"Address"`
	Username string   `yaml:"Username"`
	Password string   `yaml:"Password"`
}

type Push struct {
	WxPush map[string]string `yaml:"WxPush"`
	Email  map[string]string `yaml:"Email"`
}

type Pgsql struct {
	Dsn string `yaml:"Dsn"`
}

type Redis struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}
