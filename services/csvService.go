package services

import (
	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/LuisMG96/academy-go-q42021/repositories/inmem"
)

const csvToReadPath string = "./csvToRead.csv"

type Csv interface {
	GetAllCharacters() ([]*character.Characters, error)
	GetCharacterById(id int) (*character.Characters, error)
}

type CsvService struct {
}

func NewCsvService() Csv {
	csvService := &CsvService{}
	return csvService
}

func (csvService *CsvService) GetAllCharacters() ([]*character.Characters, error) {
	characterRepo := inmem.NewCharacterRepository()
	data, errorRe := characterRepo.FetchCharacters()
	if errorRe != nil {
		return nil, errorRe
	} else {
		return data, nil
	}
}

func (csvService *CsvService) GetCharacterById(id int) (*character.Characters, error) {
	characterRepo := inmem.NewCharacterRepository()
	data, errorRe := characterRepo.FetchCharacterById(id)
	if errorRe != nil {
		return nil, errorRe
	} else {
		return data, nil
	}
}

type empData struct {
	Name string
	Age  string
	City string
}
