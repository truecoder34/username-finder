package helper

import (
	"encoding/base64"
	"log"

	"github.com/skip2/go-qrcode"
)

type qrcodeInterface interface {
	GenerateQRCode(string, chan string)
}

type qrcodeEntity struct{}

var QRcode qrcodeInterface = &qrcodeEntity{}

func (qr *qrcodeEntity) GenerateQRCode(url string, c chan string) {
	var imageSize = 256

	qrCodeImageData, taskError := qrcode.Encode(url, qrcode.High, imageSize)
	if taskError != nil {
		log.Fatalln("Error generating QR code. ", taskError)
	}

	encodedData := base64.StdEncoding.EncodeToString(qrCodeImageData)

	c <- encodedData
}
