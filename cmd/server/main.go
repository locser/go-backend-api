package main

import "go-backend/router"

func main() {
	r := router.InitRouter()

	r.Run(":1222") // listen and serve on 0.0.0.0:1222 (for windows "localhost:1222")
}
