package main

import (
	"aparnasukesh/github.com/Go-Basic-WebApp2/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	controller.Routes(r)
	r.Run(":3000")
}
