package main

import (
	"github.com/dsolyakin/task-tracker/internal/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	http.InitRoutes(r)
	r.Run()
}
