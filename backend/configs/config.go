package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuthKey  *jwtauth.JWTAuth
}

func LoadConfig(path string) (*config, error) {
	var cfg *config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	checkError(err)
	err = viper.Unmarshal(&cfg)
	checkError(err)
	cfg.TokenAuthKey = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
