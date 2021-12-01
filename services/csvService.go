package services

import (
	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/LuisMG96/academy-go-q42021/repositories/inmem"
)

const csvToReadPath string = "./csvToRead.csv"

//Csv - Interface needs to methods GetAllCharacters and GetCharacterById
type Csv interface {
	GetAllCharacters() ([]*character.Characters, error)
	GetCharacterById(id int) (*character.Characters, error)
}

//CsvService - Struct who will containt two method implementation of Csv interface
type CsvService struct {
}

//NewCsvService - Returns a CsvService
func NewCsvService() *CsvService {
	csvService := &CsvService{}
	return csvService
}

//GetAllCharacters - Implementation of GetAllCharacters of interface Csv, use the Character Repository to get a lis of characters
func (csvService *CsvService) GetAllCharacters() ([]*character.Characters, error) {
	characterRepo := inmem.NewCharacterRepository()
	data, errorRe := characterRepo.FetchCharacters()
	if errorRe != nil {
		return nil, errorRe
	} else {
		return data, nil
	}
}

//GetCharacterById - Implementation of GetCharacterById of interface Csv, use the Character Repository to get a specific Character by Id
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
