package repository_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
}

func (suite *UserRepoTestSuite) SetupTest() {
	suite.T().Skip()
}

func (suite *UserRepoTestSuite) TestPaginate() {
	suite.T().Skip()
}

func (suite *UserRepoTestSuite) TestCreate() {
	suite.T().Skip()
}

func (suite *UserRepoTestSuite) TestFindById() {
	suite.T().Skip()
}

func (suite *UserRepoTestSuite) TestUpdate() {
	suite.T().Skip()
}

func (suite *UserRepoTestSuite) TestDelete() {
	suite.T().Skip()
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
