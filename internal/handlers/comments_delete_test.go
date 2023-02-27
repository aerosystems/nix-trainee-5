package handlers

import (
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var tableDeleteComment = testTable{
	{
		Name: "SUCCESS CASE: Delete comment by id, with JSON data",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     301,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID: 301,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"comment with ID 301 was deleted successfully"}`,
	},
	{
		Name: "SUCCESS CASE: Delete comment by id, with XML data",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     301,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID: 301,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>comment with ID 301 was deleted successfully</message></Response>`,
	},
	{
		Name: "FAIL CASE: Delete comment by id, with JSON data. Comment with id in Request param Not Found",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     301,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID: 302,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"comment with ID 302 does not exist"}`,
	},
	{
		Name: "FAIL CASE: Delete comment by id, with XML data. Comment with id in Request param Not Found",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     301,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID: 302,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>comment with ID 302 does not exist</message></Response>`,
	},
}

func (s *Suite) TestDeleteComment() {
	for _, t := range tableDeleteComment {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.Data.Comment.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
					AddRow(t.Data.Comment.ID, t.Data.Comment.PostId, t.Data.Comment.Name, t.Data.Comment.Email, t.Data.Comment.Body))
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.NewData.Comment.ID).
				WillReturnResult(sqlmock.NewResult(int64(t.NewData.Comment.ID), 1))
			s.mock.ExpectCommit()
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.NewData.Comment.ID).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(t.NewData.Comment.ID))

		if assert.NoError(s.T(), s.basehandler.DeleteComment(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
