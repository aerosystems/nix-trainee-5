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

var tableCreatePost = testTable{
	{
		Name: "SUCCESS CASE: Create Post by JSON Request Body",
		Data: DataStruct{
			Post: models.Post{},
		},
		NewData: NewDataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `{"id": 61,"userId": 7,"title": "voluptatem doloribus consectetur est ut ducimus","body": "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusCreated,
		ResponseBody:             `{"error":false,"message":"post with ID 61 was created successfully","data":{"id":61,"userId":7,"title":"voluptatem doloribus consectetur est ut ducimus","body":"ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}}`,
	},
	{
		Name: "SUCCESS CASE: Create Post by XML Request Body",
		Data: DataStruct{
			Post: models.Post{},
		},
		NewData: NewDataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `<post><id>61</id><userId>7</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></post>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusCreated,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>post with ID 61 was created successfully</message><data><id>61</id><userId>7</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></data></Response>`,
	},
	{
		Name: "FAIL CASE: Create Post by JSON Request Body, if Post is exist",
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
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `{"id": 61,"userId": 7,"title": "voluptatem doloribus consectetur est ut ducimus","body": "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}`,
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusBadRequest,
		ResponseBody:             `{"error":true,"message":"post with ID 61 exists"}`,
	},
	{
		Name: "FAIL CASE: Create Post by XML Request Body, if Post is exist",
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
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		RequestBody:              `<post><id>61</id><userId>7</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></post>`,
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusBadRequest,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>post with ID 61 exists</message></Response>`,
	},
}

func (s *Suite) TestCreatePost() {
	for _, t := range tableCreatePost {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusCreated:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.NewData.Post.ID).
				WillReturnError(gorm.ErrRecordNotFound)
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`user_id`,`title`,`body`,`id`) VALUES (?,?,?,?)")).
				WithArgs(t.NewData.Post.UserID, t.NewData.Post.Title, t.NewData.Post.Body, t.NewData.Post.ID).
				WillReturnResult(sqlmock.NewResult(int64(t.NewData.Post.ID), 1))
			s.mock.ExpectCommit()
		case http.StatusBadRequest:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.Data.Post.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
					AddRow(t.Data.Post.ID, t.Data.Post.UserID, t.Data.Post.Title, t.Data.Post.Body))
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(t.RequestBody))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(s.T(), s.basehandler.CreatePost(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
