package services

import (
	"testing"

	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCharacterRepo struct {
	mock.Mock
}

var characters = []character.Characters{
	{
		2,
		"Morty Smith",
		"Alive",
		"Human",
		"",
		"Male",
	},
	{
		7, "Abradolf Lincler", "unknown", "Human", "Genetic experiment", "Male",
	},
	{
		10, "Alan Rails", "Dead", "Human", "Superhuman (Ghost trains summoner)", "Male",
	},
	{10000, "Alan Rails", "Dead", "Human", "Superhuman (Ghost trains summoner)", "Male"},
}

func (mr mockCharacterRepo) GetAllCharacters() ([]*character.Characters, error) {
	arg := mr.Called()
	return arg.Get(0).([]*character.Characters), arg.Error(1)
}

func (mr mockCharacterRepo) GetCharacterById(id int) (*character.Characters, error) {
	arg := mr.Called(id)
	return arg.Get(0).(*character.Characters), arg.Error(1)
}

func TestCsvService_GetAllCharacters(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		hasError       bool
		error          error
	}{
		{
			"Succesfull",
			493,
			false,
			nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockCharacterRepo{}
			mock.On("GetAllCharacters").Return(tc.expectedLength, tc.error)
			service := NewCsvService()
			data, err := service.GetAllCharacters()

			assert.EqualValues(t, tc.expectedLength, len(data))
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
