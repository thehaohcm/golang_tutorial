package test

import (
	"encoding/json"
	"golang_project/models"
	"golang_project/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//1.

//2.
func TestShowFriendsByEmailSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", strings.NewReader("{\"email\":\"thehaohcm@yahoo.com.vn\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.FriendListResponse{
		Success: true,
		Friends: []string{"hao.nguyen@s3corp.com.vn"},
		Count:   1,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowFriendsByEmailEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowFriendsByEmailWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showFriendsByEmail", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

//3.
func TestShowCommonFriendListSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", strings.NewReader(`{
		"friends": [
		  "thehaohcm@yahoo.com.vn","chinh.nguyen@s3corp.com.vn"
		]
	  }`))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	exRs := models.FriendListResponse{
		Success: true,
		Friends: []string{"hao.nguyen@s3corp.com.vn"},
		Count:   1,
	}
	var modelRes models.FriendListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowCommonFriendListEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowCommonFriendListWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showCommonFriendList", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

//5.

// 6.
func TestShowSubscribingEmailListByEmailSuccessfulCode(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", strings.NewReader("{\"sender\": \"thehaohcm@yahoo.com.vn\",\"text\": \"Hello World! kate@example.com\"}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	exRs := models.GetSubscribingEmailListResponse{
		Success:    true,
		Recipients: []string{"kate@example.com"},
	}

	var modelRes models.GetSubscribingEmailListResponse
	err = json.Unmarshal(w.Body.Bytes(), &modelRes)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, exRs, modelRes)
}

func TestShowSubscribingEmailListByEmailEmptyBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestShowSubscribingEmailListByEmailWrongBody(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/friends/showSubscribingEmailListByEmail", strings.NewReader("{}"))
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", w.Body.String())
}
