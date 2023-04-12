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

	go queue.DeQueue()

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

	//차단 해놓기 (close(loop) 하기 전까지)
	loop := make(chan bool, 1)

	<-loop
}
