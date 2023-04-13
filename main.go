package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"queue/database"
	"queue/docs"
	"queue/queue"
	"time"
)

// @title queue
// @BasePath /
func main() {

	database.ConnectDB()

	r := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Logger())
	r.POST("/enqueue", queue.Enqueue)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		baseUrl := "http://localhost:8000"
		log.Printf("You can check api docs %s/swagger/index.html", baseUrl)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	go queue.DeQueue()

	loop := make(chan bool)

	<-loop
}
