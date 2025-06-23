package configurations

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Conf    Config
	ConfEnv ConfigENV
}

type Config struct {
	Port           string          `mapstructure:"port"`
	Endpoint       string          `mapstructure:"endpoint"`
	Cors           CorsConfig      `mapstructure:"cors"`
	InitialScripts []string        `mapstructure:"initial-scripts"`
	Line           LineConfig      `mapstructure:"line"`
	Originate      OriginateConfig `mapstructure:"originate"`
	Production     bool            `mapstructure:"production"`
}

type ConfigENV struct {
	Port                  string `mapstructure:"port"`
	Provider              string `mapstructure:"DB_PROVIDER"`
	Host                  string `mapstructure:"DB_HOST"`
	User                  string `mapstructure:"DB_USER"`
	Password              string `mapstructure:"DB_PASSWORD"`
	Dbname                string `mapstructure:"DB_NAME"`
	DBPort                string `mapstructure:"DB_PORT"`
	SSLMODE               string `mapstructure:"SSLMODE"`
	ConnectionMaxLifeTime int    `mapstructure:"CONNECTION_MAX_LIFE_TIME"`
	MaxIdleConns          int    `mapstructure:"MAX_IDEL_CONNS"`
	MaxOpenConns          int    `mapstructure:"MAX_OPEN_CONNS"`
	MAX_IDLE_TIME         int    `mapstructure:"MAX_IDLE_TIME"`
	Secret_token          string `mapstructure:"SECRET_TOKEN"`
	Access_token_max_age  int64  `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	SimQty                int    `mapstructure:"SIM_QTY"`
}

type CorsConfig struct {
	AllowOrigins     string `mapstructure:"allowOrigins"`
	AllowCredentials bool   `mapstructure:"allowCredentials"`
	AllowHeaders     string `mapstructure:"allowHeaders"`
	AllowMethods     string `mapstructure:"allowMethods"`
}

type LineConfig struct {
	UrlVerify  string `mapstructure:"url_verify"`
	ClientID   string `mapstructure:"client_id"`
	UrlProfile string `mapstructure:"url_profile"`
}

type OriginateConfig struct {
	Url          string `mapstructure:"url"`
	WebsocketUrl string `mapstructure:"websocket_url"`
	Timeout      int    `mapstructure:"timeout"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Extension    string `mapstructure:"extension"`
	Context      string `mapstructure:"context"`
	Priority     int    `mapstructure:"priority"`
	App          string `mapstructure:"app"`
}

func loadConfig(path, typeFIle, name string) (*Config, error) {
	conf := new(Config)
	viper.AddConfigPath(path)
	viper.SetConfigType(typeFIle)
	viper.SetConfigName(name)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func loadConfigENV(path, typeFIle, name string) (*ConfigENV, error) {
	conf := new(ConfigENV)
	viper.AddConfigPath(path)
	viper.SetConfigType(typeFIle)
	viper.SetConfigName(name)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func LoadConfigFile() (*Configs, error) {
	configs := new(Configs)
	conf, err := loadConfig("./configurations", "yaml", "config")
	if err != nil {
		return nil, err
	}

	confENV, err := loadConfigENV("./configurations", "yaml", "env")
	if err != nil {
		return nil, err
	}

	configs.Conf = *conf
	configs.ConfEnv = *confENV

	return configs, nil
}
