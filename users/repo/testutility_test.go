package repo

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var TestRepo = userRepo{}
var MockUser model.User
var Mock sqlmock.Sqlmock

const mockUserTestData = "mock_user_test_data.json"

func TestMain(m *testing.M) {
	log.Println("Setting up testing suite")
	setup()
	code := m.Run()
	log.Println("Tearing down testing suite")
	teardown()
	log.Println("Teardown successful")
	os.Exit(code)
}

func setup() {
	db, mock, err := sqlmock.NewWithDSN("postgres")
	if err != nil {
		log.Fatalf("Error not expected while initializing testing db: %v\n", err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error creating a mock database connection: %v", err)
	}

	TestRepo.db = gormDB
	Mock = mock
	readInUserTestData()
}

func teardown() {
	if err := Mock.ExpectationsWereMet(); err != nil {
		log.Printf("Expectations were not met: %v", err)
	}

}

func readInUserTestData() {
	testData, err := os.Open(mockUserTestData)
	if err != nil {
		log.Fatalf("Error opening  MockUserTestData JSON: %v", err)
	}
	defer testData.Close()

	data, err := ioutil.ReadAll(testData)
	if err != nil {
		log.Fatalf("Error reading in MockUserTestData file contents: %v", err)
	}

	if err := json.Unmarshal(data, &MockUser); err != nil {
		log.Fatalf("Error binding user struct: %v", err)
	}
	log.Println("User Mock Created Successfully.")
}
