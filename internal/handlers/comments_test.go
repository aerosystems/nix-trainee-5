package handlers

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testTable []struct {
	Name              string
	Data              models.Comment
	HeaderContentType string
	HeaderAccept      string
	StatusCode        int
	ResponseBody      string
}

var tableReadComment = testTable{
	{
		Name: "read comment by id, with JSON data",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		HeaderContentType: echo.MIMEApplicationJSON,
		HeaderAccept:      echo.MIMEApplicationJSON,
		StatusCode:        http.StatusOK,
		ResponseBody:      `{"error":false,"message":"comment with ID 301 was found successfully","data":{"id":301,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}}`,
	},
	{
		Name: "read comment by id, with XML data",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		HeaderContentType: echo.MIMEApplicationXML,
		HeaderAccept:      echo.MIMEApplicationXML,
		StatusCode:        http.StatusOK,
		ResponseBody: `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>comment with ID 301 was found successfully</message><data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data></Response>`,
	},
}

func (s *Suite) TestReadComment() {
	for _, t := range tableReadComment {
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
			WithArgs(t.Data.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
				AddRow(t.Data.ID, t.Data.PostId, t.Data.Name, t.Data.Email, t.Data.Body),
			)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.HeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.HeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(t.Data.ID))

		if assert.NoError(s.T(), s.basehandler.ReadComment(c)) {
			assert.Equal(s.T(), t.StatusCode, rec.Code)
			assert.Equal(s.T(), t.ResponseBody, strings.Replace(rec.Body.String(), "\n", "", 10))
		}
	}

}

// func (s *Suite) TestCreateComment() {
// 	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
// 		WithArgs(tt.comment.ID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
// 		AddRow(tt.comment.ID, tt.comment.PostId, tt.comment.Name, tt.comment.Email, tt.comment.Body),
// 	)

// 	s.mock.ExpectBegin()
// 	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `comments` (`post_id`,`name`,`email`,`body`,`id`) VALUES (?,?,?,?,?)")).
// 		WithArgs(tt.comment.PostId, tt.comment.Name, tt.comment.Email, tt.comment.Body, tt.comment.ID).
// 		WillReturnResult(sqlmock.NewResult(int64(tt.comment.ID), 1))
// 	s.mock.ExpectCommit()

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.JSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/comments")

// 	if assert.NoError(s.T(), s.basehandler.CreateComment(c)) {
// 		assert.Equal(s.T(), http.StatusCreated, rec.Code)
// 		assert.Equal(s.T(), tt.JSON, rec.Body.String())
// 	}
// }
