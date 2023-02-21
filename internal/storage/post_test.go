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

type SuitePosts struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *postRepo
	post       models.Post
}

func (s *SuitePosts) SetupSuite() {
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

	s.repository = NewPostRepo(s.DB)
	s.post = models.Post{
		ID:     61,
		UserID: 7,
		Title:  "voluptatem doloribus consectetur est ut ducimus",
		Body:   "ab nemo optio odio\ndelectus tenetur corporis similique nobis repellendus rerum omnis facilis\nvero blanditiis debitis in nesciunt doloribus dicta dolores\nmagnam minus velit",
	}
}

func (s *SuitePosts) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitPost(t *testing.T) {
	suite.Run(t, new(SuitePosts))
}

func (s *SuitePosts) TestFindAllPosts() {
	query := regexp.QuoteMeta("SELECT * FROM `posts`")
	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
		AddRow(s.post.ID, s.post.UserID, s.post.Title, s.post.Body)

	s.mock.ExpectQuery(query).
		WillReturnRows(rows)

	res, err := s.repository.FindAll()
	s.T().Log("!!!", err)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&[]models.Post{s.post}, res))
}

func (s *SuitePosts) TestFindByIDPost() {
	query := regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")
	rows := sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
		AddRow(s.post.ID, s.post.UserID, s.post.Title, s.post.Body)

	s.mock.ExpectQuery(query).
		WithArgs(s.post.ID).
		WillReturnRows(rows)

	res, err := s.repository.FindByID(s.post.ID)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&s.post, res))
}

func (s *SuitePosts) TestCreatePost() {
	query := regexp.QuoteMeta("INSERT INTO `posts` (`user_id`,`title`,`body`,`id`) VALUES (?,?,?,?)")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.post.UserID, s.post.Title, s.post.Body, s.post.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.post.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Create(&s.post)
	require.NoError(s.T(), err)
}

func (s *SuitePosts) TestUpdatePost() {
	query := regexp.QuoteMeta("UPDATE `posts` SET `user_id`=?,`title`=?,`body`=? WHERE `id` = ?")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.post.UserID, s.post.Title, s.post.Body, s.post.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.post.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Update(&s.post)
	require.NoError(s.T(), err)
}

func (s *SuitePosts) TestDeletePost() {
	query := regexp.QuoteMeta("DELETE FROM `posts` WHERE `posts`.`id` = ?")

	s.mock.ExpectBegin()
	s.mock.ExpectExec(query).
		WithArgs(s.post.ID).
		WillReturnResult(sqlmock.NewResult(int64(s.post.ID), 1))
	s.mock.ExpectCommit()

	err := s.repository.Delete(&s.post)
	require.NoError(s.T(), err)
}