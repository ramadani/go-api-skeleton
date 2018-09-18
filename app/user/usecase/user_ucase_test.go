package usecase_test

import (
	"testing"

	"github.com/ramadani/go-api-skeleton/app/user/usecase"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	usecase *usecase.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.usecase = usecase.New()
}

func (suite *UserUsecaseTestSuite) TestPaginate() {
	userPaginate, err := suite.usecase.Paginate(1, 10)
	suite.Nil(err)
	suite.NotEmpty(userPaginate.Data)
	suite.Equal(1, userPaginate.Page)
	suite.Equal(10, userPaginate.PerPage)
}

func (suite *UserUsecaseTestSuite) TestCreate() {
	userData, err := suite.usecase.Create("User Fullname", "foo@example.com", "randomString")
	suite.Nil(err)
	suite.NotEmpty(userData)
	suite.Equal("User Fullname", userData.Name)
	suite.Equal("foo@example.com", userData.Email)
}

func (suite *UserUsecaseTestSuite) TestFindById() {
	userData, err := suite.usecase.FindByID(1)
	suite.Nil(err)
	suite.NotEmpty(userData)
}

func (suite *UserUsecaseTestSuite) TestUpdate() {
	userData, err := suite.usecase.Update("FooBar", 1)
	suite.Nil(err)
	suite.NotEmpty(userData)
	suite.Equal("FooBar", userData.Name)
}

func (suite *UserUsecaseTestSuite) TestDelete() {
	err := suite.usecase.Delete(1)
	suite.Nil(err)
}
func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
