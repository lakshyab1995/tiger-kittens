package db

import (
	"database/sql"
	"log"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lakshyab1995/tiger-kittens/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	username = "testuser"
)

type UserSuite struct {
	suite.Suite
	DB       *gorm.DB
	mock     sqlmock.Sqlmock
	userRepo UserRepository
	mockRepo *MockUserRepository
}

func (s *UserSuite) SetupSuite() {
	var (
		err error
		db  *sql.DB
	)

	db, s.mock, err = sqlmock.New()
	assert.NoError(s.T(), err)

	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	s.userRepo = NewUserRepository(s.DB)

	s.mockRepo = NewMockUserRepository(gomock.NewController(s.T()))
}

func (s *UserSuite) BeforeTest(_, _ string) {
	if !assert.NotNil(s.T(), s.userRepo, "userRepo is not created") {
		s.T().Fatal("userRepo is not created")
	}
}

func (s *UserSuite) AfterTest(_, _ string) {
	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserSuite) TestUserRepository_Create() {
	usr := GetMockUser()

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`username`,`password`,`email`) VALUES (?,?,?,?)")).WithArgs(usr.ID, usr.Username, usr.Password, usr.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `id` FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1")).WithArgs(usr.Username).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(usr.ID))

	token, err := s.userRepo.Create(usr)

	log.Print(token)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), token)

	usrId, err := s.userRepo.GetUsrIdByUsername(usr.Username)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), usrId)
}

func (s *UserSuite) TestUserRepository_AuthenticateUser() {

}

func GetMockUser() *User {
	return &User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@email.com",
		ID:       uuid.New().String(),
	}
}

func GetMockTokenModel() *jwt.TokenModel {
	return &jwt.TokenModel{
		Token:  "testtoken",
		Expiry: "16825390",
	}
}
