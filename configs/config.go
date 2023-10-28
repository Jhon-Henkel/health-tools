package configs

import (
	"fmt"
	"os"

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
	if err != nil {
		fmt.Println("No config file loaded - using env variables")
		cfg.DbDriver = os.Getenv("DB_DRIVER")
		cfg.DbHost = os.Getenv("DB_HOST")
		cfg.DbPort = os.Getenv("DB_PORT")
		cfg.DbUser = os.Getenv("DB_USER")
		cfg.DbPassword = os.Getenv("DB_PASSWORD")
		cfg.DbName = os.Getenv("DB_NAME")
		cfg.WebServerPort = os.Getenv("WEB_SERVER_PORT")
	} else {
		err = viper.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}
	}
	cfg.TokenAuthKey = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
