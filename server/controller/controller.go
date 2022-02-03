package controller

import (
	"encoding/base64"
	"errors"
	"net/http"
	"username-finder/server/service"

	"github.com/gin-gonic/gin"

	"fmt"
	"os"

	"username-finder/go-rabbit-mq/lib/event"

	"github.com/streadway/amqp"
)

func Username(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid JSON body"))
		return
	}
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	emitter, err := event.NewEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 10; i++ {
		emitter.Push(fmt.Sprintf("[%d] - %s", i, os.Args[1]), os.Args[1])
	}

	c.JSON(http.StatusOK, matchedUrls)
}

func QRcodegenerator(c *gin.Context) {
	// 1 - get links list
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid JSON body"))
		return
	}

	// 2 - check that URL exists.
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	// generate QR codes for valif urls
	qrcodes := service.QRcodeService.QRCodeGenerate(matchedUrls)

	// 3 - send png user to picture
	qr_like_bytes, _ := base64.StdEncoding.DecodeString(qrcodes[0])

	c.Data(http.StatusOK, "image/gif", qr_like_bytes)

}
