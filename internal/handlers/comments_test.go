package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aerosystems/nix-trainee-5-6-7-8/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[int]*models.Comment{
		301: &models.Comment{
			Id:     301,
			PostId: 61,
			Name:   "quia voluptatem sunt voluptate ut ipsa",
			Email:  "Lindsey@caitlyn.net",
			Body:   "fuga aut est delectus earum optio impedit qui excepturi iusto consequatur deserunt soluta sunt et autem neque dolor ut saepe dolores assumenda ipsa eligendi",
		},
	}
	commentJSON = `{
		"postId": 61,
		"name": "quia voluptatem sunt voluptate ut ipsa",
		"email": "Lindsey@caitlyn.net",
		"body": "fuga aut est delectus earum optio impedit qui excepturi\niusto consequatur deserunt soluta sunt\net autem neque\ndolor ut saepe dolores assumenda ipsa eligendi"
	}`
	commentXML = `<id>302</id>
	<postId>61</postId>
	<name>quia voluptatem sunt voluptate ut ipsa</name>
	<email>Lindsey@caitlyn.net</email>
	<body>fuga aut est delectus earum optio impedit qui excepturi&#xA;iusto consequatur deserunt soluta sunt&#xA;et autem neque&#xA;dolor ut saepe dolores assumenda ipsa eligendi</body>`
)

func TestCreateComment(t *testing.T) {

}
