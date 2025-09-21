package main

// import (
// 	"fmt"

// 	"github.com/spf13/viper"
// )

// type Config struct {
//     Server struct {
//         Port int `mapstructure:"port"`
//     } `mapstructure:"server"`

//     Databases []struct {
//         User string `mapstructure:"user"`
//         Password string `mapstructure:"password"`
//         Host string `mapstructure:"host"`
//     }`mapstructure:"databases"`
// }

// func main() {

//     //go get github.com/spf13/viper
// 	viper := viper.New()
//     viper.AddConfigPath("./config/")

//     viper.SetConfigName("local")
//     viper.SetConfigType("yaml")

//     err := viper.ReadInConfig()
//     if err != nil {
//         panic (fmt.Errorf("failed to read configuration: %w", err))
//     }

//     fmt.Println("Server Port::", viper.GetInt("server.port"))
//     fmt.Println("Key::", viper.GetString("security.jwt.key"))

//     // config structure
//     var config Config
//     if err:= viper.Unmarshal(&config); err != nil {
//         fmt.Printf("unable to decode configuration %w", err)
//     }

//     fmt.Println("Config Port::" , config.Server.Port)

//     for _,db := range config.Databases {
//         fmt.Printf("database user: %s, pass: %s, host: %s \n", db.User, db.Password, db.Host)
//     }

// }
