package controller

import (
	"github.com/eulbyvan/go-user-management/model/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dummyUser = []entity.User{
	{
		ID:        1,
		FirstName: "first dummy 1",
		LastName:  "last dummy 1",
		Email:     "dummy1@home.com",
	},
	{
		ID:        2,
		FirstName: "first dummy 2",
		LastName:  "last dummy 2",
		Email:     "dummy2@home.com",
	},
}

type userUseCaseMock struct {
	mock.Mock
}

func (u userUseCaseMock) GetAll() ([]entity.User, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.User), nil
}

func (u userUseCaseMock) GetOne(id int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCaseMock) Add(newUser *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCaseMock) Edit(user *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCaseMock) Remove(id int) error {
	//TODO implement me
	panic("implement me")
}

type UserApiTestSuite struct {
	suite.Suite
	useCaseTest     userUseCaseMock
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *UserApiTestSuite) SetupTest() {
	suite.useCaseTest = userUseCaseMock{}
	suite.routerTest = gin.Default()
	suite.routerGroupTest = suite.routerTest.Group("/v1")
}
func (suite *UserApiTestSuite) Test_GetAllAPI_Success() {
	suite.useCaseTest.On("GetAll").Return(dummyUser, nil)
	user := NewUserController(suite.routerGroupTest, suite.useCaseTest)
	handler := user.GetAll
	suite.routerTest.GET("", handler)
	rr := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/users", nil)
	request.Header.Set("Content-Type", "application/json")

	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 200)
}
func TestUserApiTestSuite(t *testing.T) {
	suite.Run(t, new(UserApiTestSuite))
}
