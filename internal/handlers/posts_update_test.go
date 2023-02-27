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

var tableUpdatePost = testTable{
	{
		Name: "SUCCESS CASE: Update Post with JSON Request Body",
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
				ID:     61,
				UserID: 8,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `{"userId": 8,"title": "voluptatem doloribus consectetur est ut ducimus","body": "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"post with ID 61 was updated successfully"}`,
	},
	{
		Name: "SUCCESS CASE: Update Post with XML Request Body",
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
				ID:     61,
				UserID: 8,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `<post><userId>8</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></post>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>post with ID 61 was updated successfully</message></Response>`,
	},
	{
		Name: "FAIL CASE: Update Post with JSON Request Body, if Post does not exist",
		Data: DataStruct{
			Post: models.Post{
				ID:     62,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		NewData: NewDataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `{"userId": 8,"title": "voluptatem doloribus consectetur est ut ducimus","body": "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"post with ID 61 does not exists"}`,
	},
	{
		Name: "FAIL CASE: Update Post with XML Request Body, if Post does not exist",
		Data: DataStruct{
			Post: models.Post{
				ID:     62,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		NewData: NewDataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 8,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `<post><userId>8</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></post>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>post with ID 61 does not exists</message></Response>`,
	},
}

func (s *Suite) TestUpdatePost() {
	for _, t := range tableUpdatePost {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.Data.Post.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
					AddRow(t.Data.Post.ID, t.Data.Post.UserID, t.Data.Post.Title, t.Data.Post.Body))

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `user_id`=?,`title`=?,`body`=? WHERE `id` = ?")).
				WithArgs(t.NewData.Post.UserID, t.NewData.Post.Title, t.NewData.Post.Body, t.Data.Post.ID).
				WillReturnResult(sqlmock.NewResult(int64(t.NewData.Post.ID), 1))
			s.mock.ExpectCommit()
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.NewData.Post.ID).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(t.RequestBody))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/posts/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(t.NewData.Post.ID))

		if assert.NoError(s.T(), s.basehandler.UpdatePost(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
