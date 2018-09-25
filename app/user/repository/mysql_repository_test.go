package repository_test

import (
	"database/sql"
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
	repo repository.Repository
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *MySqlUserRepoTestSuite) createUserRepoInstance(repo repository.Repository) {
	suite.repo = repo
}

func (suite *MySqlUserRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	suite.Nil(err)
	suite.db = db
	suite.mock = mock
	suite.createUserRepoInstance(repository.NewMySQLRepository(db))
}

func (suite *MySqlUserRepoTestSuite) TestPaginate() {
	defer suite.db.Close()

	limit := 10
	userRows := sqlmock.NewRows([]string{"id", "name", "email"})
	totalRows := sqlmock.NewRows([]string{"total"}).AddRow(limit)

	for i := 1; i <= limit; i++ {
		userRows.AddRow(i, fmt.Sprintf("User Fullname %d", i), fmt.Sprintf("user%d@mail.com", i))
	}

	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+) OFFSET (.?) LIMIT (.?)`).
		WithArgs(0, limit).
		WillReturnRows(userRows)
	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+)`).
		WillReturnRows(totalRows)

	users, total, uErr := suite.repo.Paginate(0, uint(limit))
	suite.Nil(uErr)
	suite.Equal(uint(limit), total)
	suite.Equal(limit, len(users))
}

func (suite *MySqlUserRepoTestSuite) TestShouldCreateNewUser() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	now := time.Now().Format(format.DateTimeToString)
	suite.mock.ExpectExec(`INSERT INTO users (.+) VALUES (.+)`).
		WithArgs("FooBar", "foo@example.com", "randomstring", now, now).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

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
