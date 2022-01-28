package provider

import (
	"username-finder/server/client"
)

type checkInterface interface {
	/*
		define method checkUrl in interface to mock it with tests
	*/
	CheckUrl(string, chan string)
}

type checker struct{}

var Checker checkInterface = &checker{}

/*
	- method will be called by goroutines from service level
	- checks passed URL and sent answer to CHANNEL
	[ param 1 ] url : url
	[ param 2 ] c : chan string - channel to transport STRING data between goroutines

	reciever function.  pointer on checker struct
*/
func (check *checker) CheckUrl(url string, c chan string) {
	resp, err := client.ClientCall.GetValue(url) // call GetValue() to get data from URL

	if err != nil {
		// check if endpoint not acessable
		c <- "cant_access_resource"
		return
	}

	if resp.StatusCode > 299 {
		c <- "no_match"
	}
	if resp.StatusCode == 200 {
		c <- url
	}
}
