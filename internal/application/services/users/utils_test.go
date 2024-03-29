package users

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	models "github.com/joaofilippe/todoGo/internal/application/models/users"
)

func Test_GenerateToken(t *testing.T) {
	t.Setenv("SECRET_KEY", "secret_key")


	assert := assert.New(t)
	userUtils := new(UserUtils)

	uuid := uuid.New()

	user := models.User{
		ID:        uuid,
		FirstName: "João",
		LastName:  "Rodrigues",
		Username:  "joaofilippe",
		Email:     "joaofilippe@outlook.com",
	}

	token, err := userUtils.generateToken(user)
	assert.Nil(err)
	assert.NotEmpty(token, "token is empty")
}
