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
	//swagger 연결시켜서 channel 로 전달
	qChannel <- body.GoToQueue
}

func DeQueue() {
	//goroutine 으로 함수 실행 시켜서 for 로 계속 돌림 & channel 에 들어오는 순서대로 변수 v에 찍힘
	for v := range qChannel {
		time.Sleep(time.Second * 10)
		if err := database.DB.Create(&models.Queue{
			GoToQueue: v,
		}).Error; err != nil {
			panic(err)
		}

	}

}
