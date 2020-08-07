package users_data_transfer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserDbIsValid(t *testing.T) {

	cases := []struct {
		name   string
		userDb UserDb
		expected bool
	}{
		{
			name: "given user db with empty id when is valid then false",
			userDb:UserDb{
				Id:"",
			},
			expected: false,
		},{
			name: "given user db with id when is valid then true",
			userDb:UserDb{
				Id:"123456",
			},
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.expected, c.userDb.IsValid())
		})
	}

}
