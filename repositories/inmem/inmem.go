package inmem

import (
	"errors"
	"os"
	"sort"

	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/gocarina/gocsv"
)

//CharacterRepositoryStruct - Struct that implements CharacterRepository Interface it also contains a list of Characters
type CharacterRepositoryStruct struct {
	characters []*character.Characters
}

//FetchCharacters - implementation of FetchCharacters of CharacterRepository interface
func (c *CharacterRepositoryStruct) FetchCharacters() ([]*character.Characters, error) {
	err := getAllCharacters(c)
	if err != nil {
		return nil, err
	}
	return c.characters, nil
}

//FetchCharacterById - implementation of FetchCharacterById of CharacterRepository interface
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
		return characters[i].ID >= id
	})
	if idx < len(characters) && characters[idx].ID == id {
		return characters[idx], nil
	} else {
		return nil, errors.New("5003")
	}
}

func getAllCharacters(c *CharacterRepositoryStruct) error {
	csvFile, err := os.Open("../sample-data/characters.csv")
	if err != nil {
		return errors.New("5001")
	}
	defer csvFile.Close()
	if err := gocsv.UnmarshalFile(csvFile, &c.characters); err != nil {
		return errors.New("5002")
	}
	characterSlice := c.characters
	sort.Slice(characterSlice, func(i, j int) bool {
		return characterSlice[i].ID <= characterSlice[j].ID
	})
	c.characters = characterSlice
	return nil
}

func (c *CharacterRepositoryStruct) WriteCharactersOnCsv(characters *[]character.Characters) error {
	csvFile, err := os.Create("sample-data/charactersWrite.csv")
	if err != nil {
		return errors.New("5001")
	}
	defer csvFile.Close()
	if err = gocsv.MarshalFile(characters, csvFile); err != nil {
		return errors.New("5001")
	}
	return nil
}

//NewCharacterRepository constructor of CharacterRepositoryStruct
func NewCharacterRepository() *CharacterRepositoryStruct {
	return &CharacterRepositoryStruct{}
}
