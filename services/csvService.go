package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/LuisMG96/academy-go-q42021/repositories/inmem"
)

const csvToReadPath string = "./csvToRead.csv"
const URL_API = "https://rickandmortyapi.com/api/character"

//Csv - Interface needs to methods GetAllCharacters and GetCharacterById
type Csv interface {
	GetAllCharacters() ([]*character.Characters, error)
	GetCharacterById(id int) (*character.Characters, error)
	WriteCharactersOnCSV() error
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

func (csvService *CsvService) WriteCharactersOnCSV() error {
	characterRepo := inmem.NewCharacterRepository()
	data, errorRe := getCharactersFromAPI()
	if errorRe != nil {
		return errorRe
	}
	errorRe = characterRepo.WriteCharactersOnCsv(data.Results)
	if errorRe != nil {
		return errorRe
	} else {
		return nil
	}

}
func getCharactersFromAPI() (data *responseBody, error error) {
	resp, err := http.Get(URL_API)
	var temp responseBody

	if err != nil {
		return nil, errors.New("5004")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("5004")
	}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		return nil, errors.New("5004")

	}
	return &temp, nil
}

type responseBody struct {
	Results *[]character.Characters `json:"results"`
}
