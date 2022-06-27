package repo

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"gorm.io/gorm"
	"regexp"
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

	defer s.DB.Close()

	//When
	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

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
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).WillReturnRows(sqlmock.NewRows(nil))
	defer s.DB.Close()
	actualUsers, err := s.Repo.ListAll(context.Background())

	//Then
	s.Require().NoError(err, "error calling db for ListAll: %v", err)
	s.Require().Empty(actualUsers)
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_ListAllShouldReturnError() {
	//Given
	const userQuery = `SELECT * FROM "users"`

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).WillReturnError(errors.New("unable to return a collection of users"))
	user, err := s.Repo.ListAll(context.Background())

	//Then
	s.Require().Error(err, "error expected while retrieving all users")
	s.Require().Nil(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}
