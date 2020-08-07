package users_data_transfer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserRequest_IsValid(t *testing.T) {
	cases := []struct {
		name     string
		request  CreateUserRequest
		expected bool
	}{
		{
			name:"Create user request with empty first name",
			request: CreateUserRequest{
				FirstName: "",
				LastName: "Perez",
				Email:"juan@juan.com",
				Password:"123453677645342",
			},
			expected: false,
		},
		{
			name:"Create user request with empty last name",
			request: CreateUserRequest{
				FirstName: "Juan",
				LastName: "",
				Email:"juan@juan.com",
				Password:"123453677645342",
			},
			expected: false,
		},
		{
			name:"Create user request with empty email",
			request: CreateUserRequest{
				FirstName: "Juan",
				LastName: "Perez",
				Email:"",
				Password:"123453677645342",
			},
			expected: false,
		},
		{
			name:"Create user request with empty password",
			request: CreateUserRequest{
				FirstName: "juan",
				LastName: "Perez",
				Email:"juan@juan.com",
				Password:"",
			},
			expected: false,
		},
		{
			name:"Create user request with valid data",
			request: CreateUserRequest{
				FirstName: "juan",
				LastName: "Perez",
				Email:"juan@juan.com",
				Password:"123456745321",
			},
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.expected, c.request.IsValid())
		})
	}
}