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

type SuiteComments struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *commentRepo
	comment    models.Comment
}

func (s *SuiteComments) SetupSuite() {
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
	s.comment = models.Comment{
		ID:     302,
		PostId: 61,
		Name:   "quia voluptatem sunt voluptate ut ipsa",
		Email:  "Lindsey@caitlyn.net",
		Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
	}
}

func (s *SuiteComments) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitComments(t *testing.T) {
	suite.Run(t, new(SuiteComments))
}

func (s *SuiteComments) TestFindAllComments() {
	query := regexp.QuoteMeta("SELECT * FROM `comments`")
	rows := sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
		AddRow(s.comment.ID, s.comment.PostId, s.comment.Name, s.comment.Email, s.comment.Body)

	s.mock.ExpectQuery(query).
		WillReturnRows(rows)

	res, err := s.repository.FindAll()

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&[]models.Comment{s.comment}, res))
}

func (s *SuiteComments) TestFindByIDComment() {
	query := regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")
	rows := sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
		AddRow(s.comment.ID, s.comment.PostId, s.comment.Name, s.comment.Email, s.comment.Body)

	s.mock.ExpectQuery(query).
		WithArgs(s.comment.ID).
		WillReturnRows(rows)

	res, err := s.repository.FindByID(s.comment.ID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&s.comment, res))
}

func (s *SuiteComments) TestCreateComment() {
	query := regexp.QuoteMeta("INSERT INTO `comments` (`post_id`,`name`,`email`,`body`,`id`) VALUES (?,?,?,?,?)")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.comment.PostId, s.comment.Name, s.comment.Email, s.comment.Body, s.comment.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.comment.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Create(&s.comment)
	require.NoError(s.T(), err)
}

func (s *SuiteComments) TestUpdateComment() {
	query := regexp.QuoteMeta("UPDATE `comments` SET `post_id`=?,`name`=?,`email`=?,`body`=? WHERE `id` = ?")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.comment.PostId, s.comment.Name, s.comment.Email, s.comment.Body, s.comment.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.comment.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Update(&s.comment)
	require.NoError(s.T(), err)
}

func (s *SuiteComments) TestDeleteComment() {
	query := regexp.QuoteMeta("DELETE FROM `comments` WHERE `comments`.`id` = ?")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.comment.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.comment.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Delete(&s.comment)
	require.NoError(s.T(), err)
}
