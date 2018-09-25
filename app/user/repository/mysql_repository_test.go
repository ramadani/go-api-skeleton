package repository_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ramadani/go-api-skeleton/app/user/repository"
	"github.com/ramadani/go-api-skeleton/helpers/format"
	"github.com/stretchr/testify/suite"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type MySqlUserRepoTestSuite struct {
	suite.Suite
	repo *repository.MySQLRepository
}

func (suite *MySqlUserRepoTestSuite) TestPaginate() {
	limit := 10
	db, mock, err := sqlmock.New()
	suite.Nil(err)
	defer db.Close()

	suite.repo = repository.NewMySQLRepository(db)

	userRows := sqlmock.NewRows([]string{"id", "name", "email"})
	totalRows := sqlmock.NewRows([]string{"total"}).AddRow(limit)

	for i := 1; i <= limit; i++ {
		userRows.AddRow(i, fmt.Sprintf("User Fullname %d", i), fmt.Sprintf("user%d@mail.com", i))
	}

	mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+) OFFSET (.?) LIMIT (.?)`).
		WithArgs(0, limit).
		WillReturnRows(userRows)
	mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+)`).WillReturnRows(totalRows)

	users, total, uErr := suite.repo.Paginate(0, uint(limit))

	suite.Nil(uErr)
	suite.Equal(uint(limit), total)
	suite.Equal(limit, len(users))
}

func (suite *MySqlUserRepoTestSuite) TestShouldCreateNewUser() {
	db, mock, err := sqlmock.New()
	suite.Nil(err)
	defer db.Close()

	suite.repo = repository.NewMySQLRepository(db)

	mock.ExpectBegin()
	now := time.Now().Format(format.DateTimeToString)
	mock.ExpectExec(`INSERT INTO users (.+) VALUES (.+)`).
		WithArgs("FooBar", "foo@example.com", "randomstring", now, now).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user, err := suite.repo.Create("FooBar", "foo@example.com", "randomstring")
	suite.Nil(err)
	suite.Equal(user.Name, "FooBar")
	suite.Equal(user.Email, "foo@example.com")
}

func (suite *MySqlUserRepoTestSuite) TestFindById() {
	suite.T().Skip()
}

func (suite *MySqlUserRepoTestSuite) TestUpdate() {
	suite.T().Skip()
}

func (suite *MySqlUserRepoTestSuite) TestDelete() {
	suite.T().Skip()
}

func TestMySqlUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MySqlUserRepoTestSuite))
}
