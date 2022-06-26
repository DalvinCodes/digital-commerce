package repo

import (
	"context"
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

	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).WithArgs(

		user.ID, user.Username, user.FirstName, user.LastName,
		user.Email, user.DateOfBirth).WillReturnResult(sqlmock.NewResult(0, 1))

	if err := s.Repo.Create(context.Background(), user); err != nil {
		s.Require().Nil(err)
	}

	//Then

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}
