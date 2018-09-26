package usecase_test

import (
	"fmt"
	"testing"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/repository/mocks"
	"github.com/ramadani/go-api-skeleton/app/user/usecase"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	usecase usecase.Usecase
}

func (suite *UserUsecaseTestSuite) createUsecaseInstance(ucase usecase.Usecase) {
	suite.usecase = ucase
}

func (suite *UserUsecaseTestSuite) TestPaginate() {
	var page, offset, limit, total uint

	repo := new(mocks.Repository)
	suite.createUsecaseInstance(usecase.New(repo))

	defer repo.AssertExpectations(suite.T())

	page = 1
	offset = 0
	limit = 10
	total = 10
	users := make([]data.User, limit)

	for i := uint(0); i < limit; i++ {
		users[i] = data.User{
			ID:    uint(i + 1),
			Name:  fmt.Sprintf("FooBar %d", i),
			Email: fmt.Sprintf("foo%d@example.com", i),
		}
	}

	repo.On("Paginate", offset, limit).Return(users, limit, nil).Once()

	userPaginate, err := suite.usecase.Paginate(page, limit)
	suite.Equal(users, userPaginate.Data)
	suite.Equal(total, userPaginate.Total)
	suite.Nil(err)
}

func (suite *UserUsecaseTestSuite) TestCreate() {
	repo := new(mocks.Repository)
	suite.createUsecaseInstance(usecase.New(repo))

	defer repo.AssertExpectations(suite.T())

	id := uint(1)
	name, email, password := "FooBar", "foo@example.com", "randomstring"
	repo.On("Create", name, email, password).Return(id, nil).Once()
	repo.On("FindByID", id).Return(data.User{ID: id, Name: name, Email: email}, nil).Once()

	user, err := suite.usecase.Create(name, email, password)
	suite.Nil(err)
	suite.Equal(name, user.Name)
	suite.Equal(email, user.Email)
	suite.Equal(uint(1), user.ID)
}

func (suite *UserUsecaseTestSuite) TestFindById() {
	suite.T().Skip()
	userData, err := suite.usecase.FindByID(1)
	suite.Nil(err)
	suite.NotEmpty(userData)
}

func (suite *UserUsecaseTestSuite) TestUpdate() {
	suite.T().Skip()
	userData, err := suite.usecase.Update("FooBar", 1)
	suite.Nil(err)
	suite.NotEmpty(userData)
	suite.Equal("FooBar", userData.Name)
}

func (suite *UserUsecaseTestSuite) TestDelete() {
	suite.T().Skip()
	err := suite.usecase.Delete(1)
	suite.Nil(err)
}
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
