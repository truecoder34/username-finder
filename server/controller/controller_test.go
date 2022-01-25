package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"username-finder/server/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getUsernameService func(urls []string) []string
)

type serviceMock struct{}

func (sm *serviceMock) UsernameCheck(urls []string) []string {
	return getUsernameService(urls)
}

func TestUsername_Success(t *testing.T) {
	/*
		mock the service;
		create testing env to implement "usernameService"
	*/
	service.UsernameService = &serviceMock{}
	getUsernameService = func(urls []string) []string {
		return []string{
			"https://twitter.com/fillpackart",
			"https://github.com/truecoder34",
			"https://vk.com/gus_poet",
		}
	}
	r := gin.Default()
	jsonBody := `["https://twitter.com/fillpackart", "https://github.com/truecoder34", "https://vk.com/gus_poet",]`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 3, len(result))
}

func TestUsername_Invalid_Data(t *testing.T) {
	/*
		Specify INVALID DATA TYPE in this test. Object, not DICT
	*/
	r := gin.Default()
	//instead of using array syntax, we used object
	jsonBody := `{"https://twitter.com/fillpackart", "https://github.com/truecoder34", "https://vk.com/gus_poet"}`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}
