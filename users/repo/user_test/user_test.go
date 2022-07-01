package repo

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"gorm.io/gorm"
	"regexp"
	"time"
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	const userQuery = `INSERT INTO "users" ("id","username","first_name","last_name","email","dob","created_at","updated_at","is_deleted") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	//When
	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).
		WithArgs(
			user.ID, user.Username, user.FirstName,
			user.LastName, user.Email, user.DateOfBirth, user.CreatedAt, user.UpdatedAt, user.IsDeleted).
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
	obj := s.SeedUser()

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(obj.ID).
		WillReturnRows(sqlmock.NewRows(nil))

	//Then
	user, err := s.Repo.FindByID(context.Background(), obj.ID)
	s.Require().NoError(err, "unexpected error while creating user")
	s.Require().Empty(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_FindByID_ReturnsError() {
	//Given
	const userQuery = `SELECT * FROM "users" WHERE id = $1`
	obj := s.SeedUser()

	//When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(obj.ID).
		WillReturnError(errors.New("unable to query db for user"))

	user, err := s.Repo.FindByID(context.Background(), obj.ID)

	//Then
	s.Require().Error(err, "error was expected while retrieving user")
	s.Require().Empty(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUser_Update() {
	//Given
	const userQuery = `UPDATE "users" SET "username"=$1,"first_name"=$2,"last_name"=$3,"email"=$4,"dob"=$5,"updated_at"=$6 WHERE "id" = $7`
	user := s.SeedUser()
	user.UpdatedAt = time.Now()

	//When
	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).
		WithArgs(
			user.Username, user.FirstName, user.LastName,
			user.Email, user.DateOfBirth, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.Repo.Update(context.Background(), user)
	s.Require().Nil(err)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}
