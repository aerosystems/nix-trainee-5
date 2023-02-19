package storage

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *commentRepo
	// comment    *models.Comment
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	require.NoError(s.T(), err)

	s.repository = NewCommentRepo(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestFindAll() {
	tc := models.Comment{
		Id:     302,
		PostId: 61,
		Name:   "quia voluptatem sunt voluptate ut ipsa",
		Email:  "Lindsey@caitlyn.net",
		Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
	}

	rows := sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
		AddRow(tc.Id, tc.PostId, tc.Name, tc.Email, tc.Body)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `comments`")).
		WillReturnRows(rows)

	res, err := s.repository.FindAll()
	require.NoError(s.T(), err)

	require.Nil(s.T(), deep.Equal(&[]models.Comment{tc}, res))
}

// func TestFindAll(t *testing.T) {
// 	tc := models.Comment{
// 		Id:     302,
// 		PostId: 61,
// 		Name:   "quia voluptatem sunt voluptate ut ipsa",
// 		Email:  "Lindsey@caitlyn.net",
// 		Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
// 	}

// 	db, mock, err := sqlmock.New()
// 	require.NoError(t, err)

// 	dialector := mysql.New(mysql.Config{
// 		Conn:                      db,
// 		DriverName:                "mysql",
// 		SkipInitializeWithVersion: true,
// 	})

// 	DB, err := gorm.Open(dialector, &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Error),
// 	})
// 	require.NoError(t, err)

// 	repository := NewCommentRepo(DB)

// 	rows := sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
// 		AddRow(1, tc.PostId, tc.Name, tc.Email, tc.Body)

// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		"SELECT * FROM `comments`")).
// 		WillReturnRows(rows)

// 	res, err := repository.FindAll()
// 	require.NoError(t, err)

// 	require.Nil(t, deep.Equal(&[]models.Comment{tc}, res))
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expections: %s", err)
// 	}
// }

// func (s *Suite) TestFindByID() {
// 	var (
// 		Id     = 302
// 		PostId = 61
// 		Name   = "quia voluptatem sunt voluptate ut ipsa"
// 		Email  = "Lindsey@caitlyn.net"
// 		Body   = "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"
// 	)

// 	rows := sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
// 		AddRow(302, 61, "quia voluptatem sunt voluptate ut ipsa", "Lindsey@caitlyn.net", "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi")

// 	s.mock.ExpectBegin()
// 	s.mock.ExpectQuery(regexp.QuoteMeta(
// 		"SELECT * FROM `comments` WHERE `comments`.`id` = $1")).
// 		WithArgs(Id).
// 		WillReturnRows(rows)
// 	s.mock.ExpectCommit()

// 	res, err := s.repository.FindByID(Id)

// 	s.T().Log(Id)
// 	s.T().Log(err)
// 	s.T().Log(models.Comment{Id: Id, PostId: PostId, Name: Name, Email: Email, Body: Body})
// 	s.T().Log(res)
// 	require.NoError(s.T(), err)
// 	// require.Equal(s.T(), &models.Comment{Id: Id, PostId: PostId, Name: Name, Email: Email, Body: Body}, res, "The two words should be the same.")
// 	// require.Nil(s.T(), deep.Equal(&models.Comment{Id: Id, PostId: PostId, Name: Name, Email: Email, Body: Body}, res))

// }

// func (s *Suite) TestCreateComment() {
// 	tc := models.Comment{
// 		Id:     301,
// 		PostId: 61,
// 		Name:   "quia voluptatem sunt voluptate ut ipsa",
// 		Email:  "Lindsey@caitlyn.net",
// 		Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
// 	}
// 	s.mock.ExpectBegin()
// 	s.mock.ExpectExec(regexp.QuoteMeta(
// 		"INSERT INTO `comments` (`id`,`post_id`,`name`,`email`,`body`) VALUES ($1,$2,$3,$4,$5) RETURNING `comments`.`id`")).
// 		WithArgs(tc.Id, tc.PostId, tc.Name, tc.Email, tc.Body)
// 	s.mock.ExpectRollback()

// 	err := s.repository.Create(&tc)

// 	require.NoError(s.T(), err)
// }
