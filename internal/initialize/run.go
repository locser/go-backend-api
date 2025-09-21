package initialize

import (
	"fmt"
	"go-backend/global"
)

func Run() {

	LoadConfig()
	fmt.Println("loading configuration mysql:", global.Config.Mysql.Host)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":1222") // listen and serve on 0.0.0.0:1222 (for windows "localhost:1222")
}
