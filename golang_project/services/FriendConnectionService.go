package services

import (
	"golang_project/models"
	"golang_project/repositories"
)

type FriendConnectionService interface {
	CreateConnection(models.FriendConnectionRequest) models.FriendConnectionResponse
	GetFriendConnection(request models.FriendListRequest) models.FriendListResponse
	ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse
	SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse
	BlockSuscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse
	GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse
}

type service struct {
	repository repositories.FriendConnectionRepository
}

func New(repo repositories.FriendConnectionRepository) FriendConnectionService {
	return &service{
		repository: repo,
	}
}

func (svc *service) CreateConnection(request models.FriendConnectionRequest) models.FriendConnectionResponse {
	var response models.FriendConnectionResponse
	response.Success = svc.repository.CreateFriendConnection(request.Friends)
	return response
}

func (svc *service) GetFriendConnection(request models.FriendListRequest) models.FriendListResponse {
	var response models.FriendListResponse
	response.Friends = svc.repository.FindFriendsByEmail(request.Email)

	if response.Friends != nil && len(response.Friends) > 0 {
		response.Success = true
		response.Count = len(response.Friends)
	}
	return response
}

func (svc *service) ShowCommonFriendList(request models.CommonFriendListRequest) models.CommonFriendListResponse {
	var response models.CommonFriendListResponse
	if len(request.Friends) <= 0 {
		return response
	}
	response.Friends = svc.repository.FindCommonFriendsByEmails(request.Friends)
	if response.Friends != nil {
		response.Success = true
		response.Count = len(response.Friends)
	}
	return response
}

func (svc *service) SubscribeFromEmail(request models.SubscribeRequest) models.SubscribeResponse {
	var response models.SubscribeResponse
	response.Success = svc.repository.SubscribeFromEmail(request)
	return response
}

func (svc *service) BlockSuscribeByEmail(request models.BlockSubscribeRequest) models.BlockSubscribeResponse {
	var response models.BlockSubscribeResponse
	response.Success = svc.repository.BlockSubscribeByEmail(request)
	return response
}

func (svc *service) GetSubscribingEmailListByEmail(request models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	var response models.GetSubscribingEmailListResponse
	if request == (models.GetSubscribingEmailListRequest{}) {
		return response
	}
	response = svc.repository.GetSubscribingEmailListByEmail(request)
	return response
}
