package users_domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetFullName(t *testing.T) {

	cases := []struct {
		name   string
		user User
		expectedFullName string
	}{{
			name: "given user db with id when is valid then true",
			user : User{
				FirstName: "Juan",
				LastName: "Perez",
			},
			expectedFullName: "Perez, Juan",
	},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.expectedFullName, c.user.GetFullName())
		})
	}

}