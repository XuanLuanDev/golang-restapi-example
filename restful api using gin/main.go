package main

import (
	"gin-restful-api/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}