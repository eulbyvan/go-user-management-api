package usecase

import (
	"errors"
	"github.com/eulbyvan/go-user-management/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

type userRepoMock struct {
	mock.Mock
}

func (u userRepoMock) FindAll() ([]entity.User, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.User), nil
}

func (u userRepoMock) FindOne(id int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoMock) Create(newUser *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoMock) Update(user *entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoMock) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

type UserUseCaseTestSuite struct {
	suite.Suite
	repoMock userRepoMock
}

//Secara otomatis SetupTest dijalankan
func (suite *UserUseCaseTestSuite) SetupTest() {
	// UserUseCase membutuhkan dependency berupa UserRepo
	// Jadi kita membutuhkan mocking terhadap repo ini
	suite.repoMock = userRepoMock{}
}

func (suite *UserUseCaseTestSuite) TestUserUseCase_GetAll_Success() {
	suite.repoMock.On("FindAll").Return(dummyUser, nil)
	userUseCaseTest := NewUserUsecase(suite.repoMock)
	users, err := userUseCaseTest.GetAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(users))
}
func (suite *UserUseCaseTestSuite) TestUserUseCase_GetAll_Failed() {
	suite.repoMock.On("FindAll").Return(nil, errors.New("Failed"))
	userUseCaseTest := NewUserUsecase(suite.repoMock)
	users, err := userUseCaseTest.GetAll()
	assert.Nil(suite.T(), users)
	assert.Equal(suite.T(), "Failed", err.Error())
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
