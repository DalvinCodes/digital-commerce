package repo

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"gorm.io/gorm"
	"regexp"
)

const (
	mockID = "83297b32-a8f8-4c62-a1a6-8be574899cba"
)

func (s *UserTestSuite) TestUser_NewRepository() {

	var gormDB *gorm.DB

	got := repo.NewUserRepository(gormDB)
	want := repo.NewUserRepository(gormDB)

	s.Require().Equalf(got, want, "Got %v : Want: %v ", got, want)
}

func (s *UserTestSuite) TestUser_Create() {
	//Given
	user := s.SeedUser()
	const userQuery = `INSERT INTO "users" ("id","username","first_name","last_name","email","dob") VALUES ($1,$2,$3,$4,$5,$6)`

	//When
	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).
		WithArgs(
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(
			sqlmock.NewResult(1, 1))

	if err := s.Repo.Create(context.Background(), user); err != nil {
		s.Require().Nil(err)
	}

	//Then
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_ListAll() {
	//Given
	const userQuery = `SELECT * FROM "users"`

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WillReturnRows(sqlmock.NewRows(nil))

	actualUsers, err := s.Repo.ListAll(context.Background())

	//Then
	s.Require().NoError(err, "error calling db for ListAll: %v", err)
	s.Require().Empty(actualUsers)
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_ListAll_ReturnsError() {
	//Given
	const userQuery = `SELECT * FROM "users"`

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WillReturnError(errors.New("unable to return a collection of users"))
	user, err := s.Repo.ListAll(context.Background())

	//Then
	s.Require().Error(err, "error was expected while retrieving all users")
	s.Require().Nil(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_FindByID() {
	//Given
	const userQuery = `SELECT * FROM "users" WHERE id = $1`

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(nil))

	//Then
	user, err := s.Repo.FindByID(context.Background(), mockID)
	s.Require().NoError(err, "unexpected error while creating user")
	s.Require().Empty(user)
}

func (s *UserTestSuite) TestUser_FindByID_ReturnsError() {
	//Given
	const userQuery = `SELECT * FROM "users" WHERE id = $1`

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnError(errors.New("unable to query db for user"))
	user, err := s.Repo.FindByID(context.Background(), mockID)

	//Then
	s.Require().Error(err, "error was expected while retrieving user")
	s.Require().Empty(user)
}
