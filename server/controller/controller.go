package controller

import (
	"errors"
	"net/http"
	"username-finder/server/service"

	"github.com/gin-gonic/gin"
)

func Username(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid JSON body"))
		return
	}
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	c.JSON(http.StatusOK, matchedUrls)
}

func QRcodegenerator(c *gin.Context) {
	// 1 - get links list
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid JSON body"))
		return
	}

	// 2 - check that it exists. operate only with manthed URLS
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	// 3 - send png user to picture

	qrcode := service.QRcodeService.QRCodeGenerate(matchedUrls)

	//
	c.JSON(http.StatusOK, qrcode)
}
