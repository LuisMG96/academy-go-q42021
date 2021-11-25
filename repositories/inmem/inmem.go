package inmem

import (
	"errors"
	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/gocarina/gocsv"
	"os"
	"sort"
)

type CharacterRepositoryStruct struct {
	characters []*character.Characters
}

func (c *CharacterRepositoryStruct) FetchCharacters() ([]*character.Characters, error) {
	err := getAllCharacters(c)
	if err != nil {
		return nil, err
	}
	return c.characters, nil
}

func (c *CharacterRepositoryStruct) FetchCharacterById(id int) (*character.Characters, error) {
	err := getAllCharacters(c)
	if err != nil {
		return nil, err
	}
	data, err := searchOnSlice(c.characters, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func searchOnSlice(characters []*character.Characters, id int) (*character.Characters, error) {

	idx := sort.Search(len(characters), func(i int) bool {
		return characters[i].Id >= id
	})
	if idx < len(characters) && characters[idx].Id == id {
		return characters[idx], nil
	} else {
		return nil, errors.New("5003")
	}
}

func getAllCharacters(c *CharacterRepositoryStruct) error {
	csvFile, err := os.Open("./sample-data/characters.csv")
	if err != nil {
		return errors.New("5001")
	}
	defer csvFile.Close()
	if err := gocsv.UnmarshalFile(csvFile, &c.characters); err != nil {
		return errors.New("5002")
	}
	characterSlice := c.characters
	sort.Slice(characterSlice, func(i, j int) bool {
		return characterSlice[i].Id <= characterSlice[j].Id
	})
	c.characters = characterSlice
	return nil
}

func NewCharacterRepository() character.CharactersRepository {
	return &CharacterRepositoryStruct{}
}
