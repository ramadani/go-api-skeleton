package repository_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ramadani/go-api-skeleton/app/user/data"
	"github.com/ramadani/go-api-skeleton/app/user/repository"
	"github.com/ramadani/go-api-skeleton/commons/format"
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

	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+) LIMIT (.?) OFFSET (.?)`).
		WithArgs(limit, 0).
		WillReturnRows(userRows)
	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+)`).
		WillReturnRows(totalRows)

	users, total, err := suite.repo.Paginate(uint(limit), 0)
	suite.Nil(err)
	suite.Equal(uint(limit), total)
	suite.Equal(limit, len(users))
}

func (suite *MySqlUserRepoTestSuite) TestPaginateOnFailure() {
	defer suite.db.Close()

	limit := 10
	userRows := sqlmock.NewRows([]string{"id", "name", "email"})
	totalRows := sqlmock.NewRows([]string{"total"}).AddRow(limit)

	for i := 1; i <= limit; i++ {
		userRows.AddRow(i, fmt.Sprintf("User Fullname %d", i), fmt.Sprintf("user%d@mail.com", i))
	}

	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+) LIMIT (.?) OFFSET (.?)`).
		WithArgs(limit, 0).
		WillReturnError(fmt.Errorf("some error"))
	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+)`).
		WillReturnRows(totalRows)

	users, total, err := suite.repo.Paginate(uint(limit), 0)
	suite.NotNil(err)
	suite.Equal(uint(0), total)
	suite.Equal(0, len(users))
}

func (suite *MySqlUserRepoTestSuite) TestPaginateOnFailureGetTotal() {
	defer suite.db.Close()

	limit := 10
	userRows := sqlmock.NewRows([]string{"id", "name", "email"})

	for i := 1; i <= limit; i++ {
		userRows.AddRow(i, fmt.Sprintf("User Fullname %d", i), fmt.Sprintf("user%d@mail.com", i))
	}

	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+) LIMIT (.?) OFFSET (.?)`).
		WithArgs(limit, 0).
		WillReturnError(fmt.Errorf("some error"))
	suite.mock.ExpectQuery(`^SELECT (.+) FROM users WHERE (.+)`).
		WillReturnError(fmt.Errorf("some error"))

	users, total, err := suite.repo.Paginate(uint(limit), 0)
	suite.NotNil(err)
	suite.Equal(uint(0), total)
	suite.Equal(0, len(users))
}

func (suite *MySqlUserRepoTestSuite) TestShouldCreateNewUser() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectExec(`INSERT INTO users (.+) VALUES (.+)`).
		WithArgs("FooBar", "foo@example.com", "randomstring", now, now).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	id, err := suite.repo.Create("FooBar", "foo@example.com", "randomstring")
	suite.Nil(err)
	suite.Equal(true, id > 0)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *MySqlUserRepoTestSuite) TestShouldRollbackCreateUserOnFailure() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectExec(`INSERT INTO users (.+) VALUES (.+)`).
		WithArgs("FooBar", "foo@example.com", "randomstring", now, now).
		WillReturnError(fmt.Errorf("Some error"))
	suite.mock.ExpectRollback()

	_, err := suite.repo.Create("FooBar", "foo@example.com", "randomstring")
	suite.NotNil(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *MySqlUserRepoTestSuite) TestFindById() {
	defer suite.db.Close()

	userRows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow("1", "FooBar", "foo@example.com")

	suite.mock.ExpectQuery(`SELECT (.+) FROM users WHERE id = (.+) LIMIT 1`).
		WillReturnRows(userRows)

	user, err := suite.repo.FindByID(1)
	suite.Nil(err)
	suite.Equal(uint(1), user.ID)
	suite.Equal("FooBar", user.Name)
	suite.Equal("foo@example.com", user.Email)
}

func (suite *MySqlUserRepoTestSuite) TestShouldReturnErrorWhenFindById() {
	defer suite.db.Close()

	sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow("1", "FooBar", "foo@example.com")

	suite.mock.ExpectQuery(`SELECT (.+) FROM users WHERE id = (.+) AND deleted_at IS NULL LIMIT 1`).
		WillReturnError(fmt.Errorf("Not Found"))

	user, err := suite.repo.FindByID(2)
	suite.NotNil(err)
	suite.Equal(data.User{ID: 0, Name: "", Email: ""}, user)
}

func (suite *MySqlUserRepoTestSuite) TestShouldUpdateUser() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectExec(`UPDATE users SET (.+) WHERE id = (.) AND deleted_at IS NULL`).
		WithArgs("BarFoo", now, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	err := suite.repo.Update("BarFoo", 1)
	suite.Nil(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *MySqlUserRepoTestSuite) TestShouldRollbackUpdateUserOnFailure() {
	defer suite.db.Close()

	suite.mock.ExpectBegin()
	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectExec(`UPDATE users SET (.+) WHERE id = (.) AND deleted_at IS NULL`).
		WithArgs("BarFoo", now, 1).
		WillReturnError(fmt.Errorf("cannot update the user"))
	suite.mock.ExpectRollback()

	err := suite.repo.Update("BarFoo", 1)
	suite.NotNil(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *MySqlUserRepoTestSuite) TestShouldDeleteUser() {
	defer suite.db.Close()

	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`UPDATE users SET deleted_at = (.) WHERE id = (.) AND deleted_at IS NULL`).
		WithArgs(now, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	err := suite.repo.Delete(1)
	suite.Nil(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *MySqlUserRepoTestSuite) TestShouldRollbackDeleteUserOnFailure() {
	defer suite.db.Close()

	now := time.Now().UTC().Format(format.DateTimeToString)
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(`UPDATE users SET deleted_at = (.) WHERE id = (.) AND deleted_at IS NULL`).
		WithArgs(now, 1).
		WillReturnError(fmt.Errorf("some error"))
	suite.mock.ExpectRollback()

	err := suite.repo.Delete(1)
	suite.NotNil(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Nil(err)
}

func TestMySqlUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(MySqlUserRepoTestSuite))
}
