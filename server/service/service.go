package service

import (
	"username-finder/server/provider"
)

type usernameCheck struct{}
type qrcodeGenerate struct{}

type usernameService interface {
	UsernameCheck(urls []string) []string
}

type qrcodeService interface {
	QRCodeGenerate(urls []string) []string
}

var (
	/*
		usernameCheck struct implements usernameService interface
	*/
	UsernameService usernameService = &usernameCheck{}

	QRcodeService qrcodeService = &qrcodeGenerate{}
)

/*
	[ param 1 ] - recieve slice of URLs to process
*/
func (u *usernameCheck) UsernameCheck(urls []string) []string {
	c := make(chan string) // create channel via make
	var links []string
	matchingLinks := []string{}

	for _, url := range urls {
		/*
			start up goroutine for each url
			any data\error obtained ---> to channel
		*/
		go provider.Checker.CheckUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		/*
			get value for each url from channel and put inside slice
		*/
		links = append(links, <-c)
	}

	/*
		Remove the "no_match" and "cant_access_resource" values from the links array. Kind of Filtering
	*/
	for _, v := range links {
		if v == "cant_access_resource" {
			continue
		}
		if v == "no_match" {
			continue
		}
		matchingLinks = append(matchingLinks, v)
	}
	return matchingLinks
}

/*
	Method should return QR code object
*/
func (u *qrcodeGenerate) QRCodeGenerate(urls []string) []string {

	qrCodeObject := []string{}

	return qrCodeObject
}
