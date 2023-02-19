package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mockDB = models.Comment{
	Id:     301,
	PostId: 61,
	Name:   "quia voluptatem sunt voluptate ut ipsa",
	Email:  "Lindsey@caitlyn.net",
	Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
}
var commentJSON = `{
		"postId": 61,
		"name": "quia voluptatem sunt voluptate ut ipsa",
		"email": "Lindsey@caitlyn.net",
		"body": "fuga aut est delectus earum optio impedit qui excepturi\niusto consequatur deserunt soluta sunt\net autem neque\ndolor ut saepe dolores assumenda ipsa eligendi"
	}`

// var commentXML = `<id>302</id>
// 	<postId>61</postId>
// 	<name>quia voluptatem sunt voluptate ut ipsa</name>
// 	<email>Lindsey@caitlyn.net</email>
// 	<body>fuga aut est delectus earum optio impedit qui excepturi&#xA;iusto consequatur deserunt soluta sunt&#xA;et autem neque&#xA;dolor ut saepe dolores assumenda ipsa eligendi</body>`

// type addCommentFakeRepo struct {
// 	MockAddFunc func(models.Comment) error
// }

// func (fake *addCommentFakeRepo) Add(c models.Comment) error {
// 	return fake.MockAddFunc(c)
// }

// func newAddCommentFakeRepo() *addCommentFakeRepo {
// 	return &addCommentFakeRepo{
// 		MockAddFunc: func(c models.Comment) error { return nil },
// 	}
// }

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository models.CommentRepository
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
		Logger: logger.Default.LogMode(logger.Info),
	})
	require.NoError(s.T(), err)

	s.repository = storage.NewCommentRepo(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreateComment() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(commentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/comments/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	h := BaseHandler{commentRepo: s.repository}

	if assert.NoError(s.T(), h.CreateComment(c)) {
		assert.Equal(s.T(), http.StatusCreated, rec.Code)
		assert.Equal(s.T(), commentJSON, rec.Body.String())
	}
}
