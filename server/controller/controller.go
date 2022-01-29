package controller

import (
	"encoding/base64"
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

	// 2 - check that URL exists.
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	// generate QR codes for valif urls
	qrcodes := service.QRcodeService.QRCodeGenerate(matchedUrls)

	// 3 - send png user to picture
	qr_like_bytes, _ := base64.StdEncoding.DecodeString(qrcodes[0])

	c.Data(http.StatusOK, "image/gif", qr_like_bytes)

}
