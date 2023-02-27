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

var tableReadPosts = testTable{
	{
		Name: "SUCCESS CASE: Read all posts, with JSON data",
		Data: DataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		NewData: NewDataStruct{
			Post: models.Post{},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `{"error":false,"message":"all posts with ID were found successfully","data":[{"id":61,"userId":7,"title":"voluptatem doloribus consectetur est ut ducimus","body":"ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}]}`,
	},
	{
		Name: "SUCCESS CASE: Read all posts, with XML data",
		Data: DataStruct{
			Post: models.Post{
				ID:     61,
				UserID: 7,
				Title:  "voluptatem doloribus consectetur est ut ducimus",
				Body:   "ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit",
			},
		},
		NewData: NewDataStruct{
			Post: models.Post{},
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>all posts with ID were found successfully</message><data><id>61</id><userId>7</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></data></Response>`,
	},
	{
		Name: "FAIL CASE: Read all posts, with JSON data, if posts do not exist",
		Data: DataStruct{
			Post: models.Post{},
		},
		NewData: NewDataStruct{
			Post: models.Post{},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"posts do not exist"}`,
	},
	{
		Name: "FAIL CASE: Read all posts, with XML data, if posts do not exist",
		Data: DataStruct{
			Post: models.Post{},
		},
		NewData: NewDataStruct{
			Post: models.Post{},
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>posts do not exist</message></Response>`,
	},
}

var tableReadPost = testTable{
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
		ResponseBody:             `{"error":false,"message":"post with ID 61 was found successfully","data":{"id":61,"userId":7,"title":"voluptatem doloribus consectetur est ut ducimus","body":"ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit"}}`,
	},
	{
		Name: "SUCCESS CASE: Read post by id, with XML data",
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
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusOK,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>false</error><message>post with ID 61 was found successfully</message><data><id>61</id><userId>7</userId><title>voluptatem doloribus consectetur est ut ducimus</title><body>ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit</body></data></Response>`,
	},
	{
		Name: "FAIL CASE: Read post by id, with JSON data, if post does exist",
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
				ID: 61,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationJSON,
		RequestHeaderAccept:      echo.MIMEApplicationJSON,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `{"error":true,"message":"post with ID 61 does not exist"}`,
	},
	{
		Name: "FAIL CASE: Read post by id, with XML data, if post does exist",
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
				ID: 61,
			},
		},
		RequestHeaderContentType: echo.MIMEApplicationXML,
		RequestHeaderAccept:      echo.MIMEApplicationXML,
		ResponseStatusCode:       http.StatusNotFound,
		ResponseBody:             `<?xml version="1.0" encoding="UTF-8"?><Response><error>true</error><message>post with ID 61 does not exist</message></Response>`,
	},
}

func (s *Suite) TestReadPosts() {
	for _, t := range tableReadPosts {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts`")).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
					AddRow(t.Data.Post.ID, t.Data.Post.UserID, t.Data.Post.Title, t.Data.Post.Body))
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts`")).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/posts", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(s.T(), s.basehandler.ReadPosts(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}

func (s *Suite) TestReadPost() {
	for _, t := range tableReadPost {
		s.T().Log(t.Name)

		switch t.ResponseStatusCode {
		case http.StatusOK:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.Data.Post.ID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "title", "body"}).
					AddRow(t.Data.Post.ID, t.Data.Post.UserID, t.Data.Post.Title, t.Data.Post.Body))
		case http.StatusNotFound:
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
				WithArgs(t.NewData.Post.ID).
				WillReturnError(gorm.ErrRecordNotFound)
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, t.RequestHeaderContentType)
		req.Header.Set(echo.HeaderAccept, t.RequestHeaderAccept)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/posts/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(t.NewData.Post.ID))

		if assert.NoError(s.T(), s.basehandler.ReadPost(c)) {
			assert.Equal(s.T(), t.ResponseStatusCode, rec.Code)
			assert.Equal(s.T(), strings.Replace(t.ResponseBody, "\n", "", -1), strings.Replace(rec.Body.String(), "\n", "", -1))
		}
	}
}
