package controllers

import (
	"golang_project/models"
	"golang_project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FriendConnectionController interface {
	CreateFriendConnection(c *gin.Context)
	GetFriendListByEmail(c *gin.Context)
	ShowCommonFriendList(c *gin.Context)
	SubscribeFromEmail(c *gin.Context)
	BlockSuscribeByEmail(c *gin.Context)
	GetSubscribingEmailListByEmail(c *gin.Context)
}

type controller struct {
	service services.FriendConnectionService
}

func New(service services.FriendConnectionService) FriendConnectionController {
	return &controller{
		service: service,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 1: As a user, I need an API to create a friend connection between two email addresses.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.FriendConnectionRequest true "Create a friend connection between 2 user email"
// @Router /friends/createConnection [post]
func (ctl *controller) CreateFriendConnection(c *gin.Context) {
	var newFriendConnection models.FriendConnectionRequest
	if err := c.BindJSON(&newFriendConnection); err != nil {
		panic(err)
	}

	if len(newFriendConnection.Friends) != 2 {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}

	response := ctl.service.CreateConnection(newFriendConnection)
	c.IndentedJSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 2: As a user, I need an API to retrieve the friends list for an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.FriendListRequest true "Get a list of friend by user email"
// @Router /friends/showFriendsByEmail [post]
func (ctl *controller) GetFriendListByEmail(c *gin.Context) {
	var request models.FriendListRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if request == (models.FriendListRequest{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	response := ctl.service.GetFriendConnection(request)
	c.IndentedJSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 3: As a user, I need an API to retrieve the common friends list between two email addresses.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.CommonFriendListRequest true "Retrieve the common friends list between two email addresses"
// @Router /friends/showCommonFriendList [post]
func (ctl *controller) ShowCommonFriendList(c *gin.Context) {
	var request models.CommonFriendListRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if request.Friends == nil || len(request.Friends) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	response := ctl.service.ShowCommonFriendList(request)
	c.IndentedJSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 4: As a user, I need an API to subscribe to updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.SubscribeRequest true "Subscribe to updates from an email address"
// @Router /friends/subscribeUpdateByEmail [post]
func (ctl *controller) SubscribeFromEmail(c *gin.Context) {
	var request models.SubscribeRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	response := ctl.service.SubscribeFromEmail(request)
	c.IndentedJSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 5: As a user, I need an API to block updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.BlockSubscribeRequest true "Block updates from an email address"
// @Router /friends/blockSubscribeUpdateByEmail [post]
func (ctl *controller) BlockSuscribeByEmail(c *gin.Context) {
	var request models.BlockSubscribeRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	response := ctl.service.BlockSuscribeByEmail(request)
	c.IndentedJSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description Requirement 6: As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
// @Tags Friend API
// @Accept json
// @Produce json
// @Param   Request body models.GetSubscribingEmailListRequest true "retrieve all email addresses that can receive update from an email address"
// @Router /friends/showSubscribingEmailListByEmail [post]
func (ctl *controller) GetSubscribingEmailListByEmail(c *gin.Context) {
	var request models.GetSubscribingEmailListRequest
	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if request == (models.GetSubscribingEmailListRequest{}) {
		c.Status(http.StatusBadRequest)
		return
	}
	response := ctl.service.GetSubscribingEmailListByEmail(request)
	c.IndentedJSON(http.StatusOK, response)
}
