package queue

import (
	"github.com/gin-gonic/gin"
	"queue/database"
	"queue/models"
	"sync"
	"time"
)

type Queue struct {
	GoToQueue string `json:"go_to_queue" binding:"required"`
}

var qChannel = make(chan string, 100)

var Wg = sync.WaitGroup{}

// @tags queue
// @Summary queue
// @Description 큐 삽입
// @Accept json
// @produce json
// @Param body body queue.Queue true "데이터"
// @Success 200
// @Failure 400
// @Router /enqueue [POST]
func Enqueue(c *gin.Context) {
	var body Queue
	if err := c.ShouldBind(&body); err != nil {
		panic(err)
	}

	qChannel <- body.GoToQueue
}

func DeQueue() {

	for v := range qChannel {
		time.Sleep(time.Second * 10)
		if err := database.DB.Create(&models.Queue{
			GoToQueue: v,
		}).Error; err != nil {
			panic(err)
		}

	}

}
