package initialize

import (
	"fmt"
	"go-backend/global"

	"github.com/spf13/viper"
)

func LoadConfig() {

	//go get github.com/spf13/viper
	viper := viper.New()
	viper.AddConfigPath("./config/")

	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Key::", viper.GetString("security.jwt.key"))

	// config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %w", err)
	}

	// fmt.Println("Config Port::", config.Server.Port)

	// for _, db := range config.Mysql {
	// 	fmt.Printf("database user: %s, pass: %s, host: %s \n", db.User, db.Password, db.Host)
	// }

}
