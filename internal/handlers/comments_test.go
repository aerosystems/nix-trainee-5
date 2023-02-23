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
)

type testTable []struct {
	Name                     string
	Data                     models.Comment
	RequestParam             map[string]string
	RequestBody              string
	RequestHeaderContentType string
	RequestHeaderAccept      string
	ResponseStatusCode       int
	ResponseBody             string
}

var tableReadComment = testTable{
	{
		Name: "SUCCESS: Read comment by id, with JSON data",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestParam: map[string]string{
			"id": "301",
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"comment with ID 301 was found successfully","data":{"id":301,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}}`,
	},
	{
		Name: "SUCCESS: Read comment by id, with XML data",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestParam: map[string]string{
			"id": "301",
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>comment with ID 301 was found successfully</message><data><id>301</id><postId>61</postId><name>quia voluptatem sunt voluptate ut ipsa</name><email>Lindsey@caitlyn.net</email><body>fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi</body></data></Response>`,
	},
	{
		Name: "ERROR: Read comment by id, with JSON data. Comment with id in Request param Not Found",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestParam: map[string]string{
			"id": "302",
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"comment with ID 302 does not exist"}`,
	},
	{
		Name: "ERROR: Read comment by id, with XML data. Comment with id in Request param Not Found",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestParam: map[string]string{
			"id": "302",
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>comment with ID 302 does not exist</message></Response>`,
	},
}

func (s *Suite) TestReadComment() {
	for _, t := range tableReadComment {
		s.T().Log(t.Name)
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
			WithArgs(t.Data.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "name", "email", "body"}).
				AddRow(t.Data.ID, t.Data.PostId, t.Data.Name, t.Data.Email, t.Data.Body),
			)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/comments/:id")
		c.SetParamNames("id")
		c.SetParamValues(t.RequestParam["id"])

		if assert.NoError(s.T(), s.basehandler.ReadComment(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}

var tableCreateComment = testTable{
	{
		Name: "SUCCESS: Create comment by Request Body",
		Data: models.Comment{
			ID:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
		RequestParam: map[string]string{
			"id": "302",
		},
		RequestBody:              `{"id":301,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusCreated,
		ResponseBody:             `{"error":false,"message":"comment with ID 301 was found successfully","data":{"id":301,"postId":61,"name":"quia voluptatem sunt voluptate ut ipsa","email":"Lindsey@caitlyn.net","body":"fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi"}}`,
	},
}

func (s *Suite) TestCreateComment() {
	for _, t := range tableCreateComment {
		s.T().Log(t.Name)

		s.mock.ExpectBegin()
		s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `comments` WHERE `comments`.`id` = ?")).
			WithArgs(t.Data.ID)

		s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `comments` (`post_id`,`name`,`email`,`body`,`id`) VALUES (?,?,?,?,?)")).
			WithArgs(t.Data.PostId, t.Data.Name, t.Data.Email, t.Data.Body, t.Data.ID).
			WillReturnResult(sqlmock.NewResult(int64(t.Data.ID), 1))
		s.mock.ExpectCommit()

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
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.JSON))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// c.SetPath("/comments")

	// if assert.NoError(s.T(), s.basehandler.CreateComment(c)) {
	// 	assert.Equal(s.T(), http.StatusCreated, rec.Code)
	// 	assert.Equal(s.T(), tt.JSON, rec.Body.String())
	// }
}
