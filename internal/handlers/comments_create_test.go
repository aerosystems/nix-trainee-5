package handlers

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var tableCreateComment = testTable{
	{
		Name: "SUCCESS CASE: Create Comment by JSON Request Body",
		Data: DataStruct{
			Comment: models.Comment{},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		RequestBody:              `{"id":302,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusCreated,
		ResponseBody:             `{"error":false,"message":"comment with ID 302 was created successfully","data":{"id":302,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}}`,
	},
	{
		Name: "SUCCESS CASE: Create Comment by XML Request Body",
		Data: DataStruct{
			Comment: models.Comment{},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		RequestBody:              `<data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusCreated,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>comment with ID 302 was created successfully</message><data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data></Response>`,
	},
	{
		Name: "FAIL CASE: Create Comment by JSON Request Body, if Comment already exists",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		RequestBody:              `{"id":302,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusBadRequest,
		ResponseBody:             `{"error":true,"message":"comment with ID 302 exists"}`,
	},
	{
		Name: "FAIL CASE: Create Comment by XML Request Body, if Comment already exists",
		Data: DataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		NewData: NewDataStruct{
			Comment: models.Comment{
				ID:     302,
				PostId: 61,
				Name:   "quia voluptatem sunt voluptate ut ipsa",
				Email:  "Lindsey@caitlyn.net",
				Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
			},
		},
		RequestBody:              `<data><id>302</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusBadRequest,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>comment with ID 302 exists</message></Response>`,
	},
}

func (s *Suite) TestCreateComment() {
	for _, t := range tableCreateComment {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusCreated:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.NewData.Comment.ID).
				WillReturnError(gorm.ErrRecordNotFound)
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `comments` (`post_id`,`name`,`email`,`body`,`id`) VALUES (?,?,?,?,?)")).
				WithArgs(t.NewData.Comment.PostId, t.NewData.Comment.Name, t.NewData.Comment.Email, t.NewData.Comment.Body, t.NewData.Comment.ID).
				WillReturnResult(sqlmock.NewResult(int64(t.NewData.Comment.ID), 1))
			s.mock.ExpectCommit()
		case http.StatusBadRequest:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
				WithArgs(t.Data.Comment.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
					AddRow(t.Data.Comment.ID, t.Data.Comment.PostId, t.Data.Comment.Name, t.Data.Comment.Email, t.Data.Comment.Body))
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/comments", strings.NewReader(t.RequestBody))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(s.T(), s.basehandler.CreateComment(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
