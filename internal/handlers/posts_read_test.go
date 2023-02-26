package handlers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
)

var tableReadPosts = testTable{
	{
		Name: "SUCCESS CASE: Read post by id, with JSON data",
		Data: DataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		NewData: NewDataStruct{
			Post: models.Post{
				ID: 61,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"comment with ID 301 was found successfully","data":{"id":301,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}}`,
	},
}

func (s *Suite) TestReadPosts() {
	for _, t := range tableReadPosts {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts`")).
				WithArgs(t.Data.Post.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
					AddRow(t.Data.Post.ID, t.Data.Post.UserID, t.Data.Post.Title, t.Data.Post.Body))
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts`")).
				WithArgs(t.NewData.Post.ID).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/posts", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(s.T(), s.basehandler.ReadPost(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
