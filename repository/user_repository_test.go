package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eulbyvan/go-user-management/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
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

type UserRepositoryTestSuite struct {
	suite.Suite
	mockResource *sql.DB
	mock         sqlmock.Sqlmock
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.mockResource.Close()

}
func (suite *UserRepositoryTestSuite) SetupTest() {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	suite.mockResource = mockDb
	suite.mock = mock
}

func (suite *UserRepositoryTestSuite) TestUserRepo_FindAll_Success() {
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"})
	for _, d := range dummyUser {
		rows.AddRow(d.ID, d.FirstName, d.LastName, d.Email)
	}
	suite.mock.ExpectQuery("SELECT (.+) FROM users").
		WillReturnRows(rows)

	repo := NewUserRepository(suite.mockResource)
	all, err := repo.FindAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))
	assert.Equal(suite.T(), 1, all[0].ID)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
