package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Db  Db
	App App
	//Jwt   *JwtConfig
}

type Db struct {
	DbUrl            string
	DbMaxConnections int
}

type App struct {
	ServerPort int
	LogLevel   string
	ImageUrl   string
}

//type JwtConfig struct {
//	AccessTokenSecret       string
//	RefreshTokenSecret      string
//	VerificationTokenSecret string
//	AccessTokenExpiry       time.Duration
//	RefreshTokenExpiry      time.Duration
//	ContextKey              string
//}

var appConfig = Config{}

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read env file: %v", err)
	}
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
		viper.SetConfigFile("base.env")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	appConfig = Config{
		Db: Db{
			DbUrl:            viper.GetString("DB_URL"),
			DbMaxConnections: viper.GetInt("DB_MAX_CONNECTIONS")},
		App: App{
			LogLevel:   viper.GetString("LogLevel"),
			ServerPort: viper.GetInt("SERVER_PORT"),
			ImageUrl:   viper.GetString("IMAGE_URL"),
		},
		//Jwt: &JwtConfig{
		//	AccessTokenSecret:  "accesstokensecret",
		//	RefreshTokenSecret: "refreshtokensecret",
		//	AccessTokenExpiry:  24,
		//	RefreshTokenExpiry: 168,
		//	ContextKey:         "user",
		//},

	}
}

func GetConfig() Config {
	return appConfig
}
