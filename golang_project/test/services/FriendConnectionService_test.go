package test

import (
	"golang_project/controllers"
	"golang_project/models"
	"golang_project/repositories"
	"golang_project/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	friendConnectionRepository repositories.FriendConnectionRepository = repositories.New()
	friendConnectionService    services.FriendConnectionService        = services.New(friendConnectionRepository)
	friendConnectionController controllers.FriendConnectionController  = controllers.New(friendConnectionService)
)

//2.
func TestShowFriendsByEmailSuccessfulCase(t *testing.T) {

	request := models.FriendListRequest{
		Email: "thehaohcm@yahoo.com.vn",
	}

	response := friendConnectionService.GetFriendConnection(request)

	exp := models.FriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
}

func TestShowFriendsByEmailEmptyModel(t *testing.T) {

	response := friendConnectionService.GetFriendConnection(models.FriendListRequest{})

	exp := models.FriendListResponse{}
	assert.Equal(t, exp, response)
}

//3.
func TestShowCommonFriendListSuccessfulCase(t *testing.T) {

	request := models.CommonFriendListRequest{
		Friends: []string{"thehaohcm@yahoo.com.vn", "chinh.nguyen@s3corp.com.vn"},
	}

	response := friendConnectionService.ShowCommonFriendList(request)

	exp := models.CommonFriendListResponse{
		Success: true,
		Friends: []string{
			"hao.nguyen@s3corp.com.vn",
		},
		Count: 1,
	}
	assert.Equal(t, exp, response)
}

func TestShowCommonFriendListEmptyModel(t *testing.T) {

	response := friendConnectionService.ShowCommonFriendList(models.CommonFriendListRequest{})

	exp := models.CommonFriendListResponse{}
	assert.Equal(t, exp, response)
}

//6.
func TestGetSubscribingEmailListByEmailSuccessfulCase(t *testing.T) {

	model := models.GetSubscribingEmailListRequest{
		Sender: "thehaohcm@yahoo.com.vn",
		Text:   "helloworld! kate@example.com",
	}

	response := friendConnectionService.GetSubscribingEmailListByEmail(model)

	exp := models.GetSubscribingEmailListResponse{
		Success: true,
		Recipients: []string{
			"kate@example.com",
		},
	}
	assert.Equal(t, exp, response)
}

func TestGetSubscribingEmailListByEmailEmptyModel(t *testing.T) {

	response := friendConnectionService.GetSubscribingEmailListByEmail(models.GetSubscribingEmailListRequest{})

	exp := models.GetSubscribingEmailListResponse{}
	assert.Equal(t, exp, response)
}

// type Stringer interface {
// 	String() string
// }

// type SendFunc struct {
// 	mock.Mock
// }

// func testMethod() {
// 	requesterMock := mocks.NewRequester(t)
// 	requesterMock.EXPECT().Get("some path").Return("result", nil)
// 	requesterMock.EXPECT().
// 		Get(mock.Anything).
// 		Run(func(path string) { fmt.Println(path, "was called") }).
// 		// Can still use return functions by getting the embedded mock.Call
// 		Call.Return(func(path string) string { return "result for " + path }, nil)
// }

// func (m *Stringer) String() string {
// 	ret := m.Called()

// 	var r0 string
// 	if rf, ok := ret.Get(0).(func() string); ok {
// 		r0 = rf()
// 	} else {
// 		r0 = ret.Get(0).(string)
// 	}

// 	return r0
// }

// // NewStringer creates a new instance of Stringer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// func NewStringer(t testing.TB) *Stringer {
// 	mock := &Stringer{}
// 	mock.Mock.Test(t)

// 	t.Cleanup(func() { mock.AssertExpectations(t) })

// 	return mock
// }
