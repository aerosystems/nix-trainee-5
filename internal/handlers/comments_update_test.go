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
	"strconv"
	"strings"
)

var tableUpdateComment = testTable{
	{
		Name: "SUCCESS CASE: Update Comment by JSON Request Body",
		Data: models.Comment{
			ID:     302,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestData: models.Comment{
			ID:     302,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestBody:              `{"id":302,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"comment with ID 302 was updated successfully"}`,
	},
	{
		Name: "SUCCESS CASE: Update Comment by XML Request Body",
		Data: models.Comment{
			ID:     302,
			PostId: 62,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestData: models.Comment{
			ID:     302,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestBody:              `<data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>comment with ID 302 was updated successfully</message></Response>`,
	},
	{
		Name: "FAIL CASE: Update Comment by JSON Request Body, if Comment does not exist",
		Data: models.Comment{},
		RequestData: models.Comment{
			ID:     302,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestBody:              `{"id":302,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"comment with ID 302 does not exists"}`,
	},
	{
		Name: "FAIL CASE: Update Comment by XML Request Body, if Comment does not exist",
		Data: models.Comment{},
		RequestData: models.Comment{
			ID:     302,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestBody:              `<data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>comment with ID 302 does not exists</message></Response>`,
	},
}

func (s *Suite) TestUpdateComment() {
	for _, t := range tableUpdateComment {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.Data.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
					AddRow(t.Data.ID, t.Data.PostId, t.Data.Name, t.Data.Email, t.Data.Body))

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `comments` SET `post_id`=?,`name`=?,`email`=?,`body`=? WHERE `id` = ?")).
				WithArgs(t.RequestData.PostId, t.RequestData.Name, t.RequestData.Email, t.RequestData.Body, t.Data.ID).
				WillReturnResult(sqlmock.NewResult(int64(t.RequestData.ID), 1))
			s.mock.ExpectCommit()
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.RequestData.ID).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(t.RequestBody))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(t.RequestData.ID))

		if assert.NoError(s.T(), s.basehandler.UpdateComment(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
