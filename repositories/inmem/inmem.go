package inmem

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/LuisMG96/academy-go-q42021/common"
	character "github.com/LuisMG96/academy-go-q42021/repositories/characters"
	"github.com/gocarina/gocsv"
)

const odd = "odd"
const charactesCsvFile = "./sample-data/characters.csv"
const charactesCsvFileWrite = "./sample-data/charactersWrite.csv"

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
	csvFile, err := os.Open(charactesCsvFile)
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

//WriteCharactersOnCsv - Function thar receive a slice of characters and write it on a CSV file
func (c *CharacterRepositoryStruct) WriteCharactersOnCsv(characters *[]character.Characters) error {
	csvFile, err := os.Create(charactesCsvFileWrite)
	if err != nil {
		return errors.New("5001")
	}
	defer csvFile.Close()
	if err = gocsv.MarshalFile(characters, csvFile); err != nil {
		return errors.New("5001")
	}
	return nil
}

func (c *CharacterRepositoryStruct) ReadWithWorkerPool(filter *common.Filter) ([]*character.Characters, error) {
	var numWps float64
	csvFile, err := os.Open(charactesCsvFile)
	if err != nil {
		return nil, errors.New("500")
	}

	fcsv := csv.NewReader(csvFile)
	rs := make([]*character.Characters, 0)
	if filter.Items == -1 {
		numWps = 100
	} else {
		if filter.Items < filter.ItemsPerWorker {
			numWps = 1
		} else {
			numWps = math.Floor(float64(filter.Items / filter.ItemsPerWorker))
		}
	}
	jobs := make(chan []string)
	res := createChanCharactes(filter)
	var wg sync.WaitGroup

	// init workers
	for w := 0; w < int(numWps); w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workersReadCsvfunc(jobs, res, filter)
		}()
	}

	go readFileConcurrently(fcsv, jobs)
	go waitAndClose(&wg, res)
	rs = addToRes(res, rs)
	fmt.Println("Longitud de res", len(rs))
	return rs, nil
}
func createChanCharactes(filter *common.Filter) chan *character.Characters {
	if filter.Items != -1 && filter.Items != 0 {
		return make(chan *character.Characters, int(filter.Items))
	} else {
		return make(chan *character.Characters)
	}
}
func workersReadCsvfunc(jobs <-chan []string, results chan<- *character.Characters, filter *common.Filter) {
	var numItem int
	if filter.Items < filter.ItemsPerWorker {
		numItem = int(filter.Items)
	} else {
		numItem = int(filter.ItemsPerWorker)
	}
	if filter.Items == -1 {
		for {
			job, ok := <-jobs
			if !ok {
				return
			}
			newCharacter := parseStruct(job)
			if filter.TypeFilter != "" {
				if filter.TypeFilter == odd {
					if odd := newCharacter.ID % 2; odd != 0 {
						results <- newCharacter
					}
				} else {
					if even := newCharacter.ID % 2; even == 0 {
						results <- newCharacter
					}
				}
			} else {
				results <- newCharacter
			}
		}
	} else {
		for i := 0; i < numItem; i++ {
			job, ok := <-jobs
			if !ok {
				return
			}
			newCharacter := parseStruct(job)
			if filter.TypeFilter != "" {
				if filter.TypeFilter == odd {
					if odd := newCharacter.ID % 2; odd != 0 {
						results <- newCharacter
					} else {
						numItem++
					}
				} else {
					if even := newCharacter.ID % 2; even == 0 {
						results <- newCharacter
					} else {
						numItem++
					}
				}
			} else {
				results <- newCharacter
			}
		}
	}
}

func readFileConcurrently(fcsv *csv.Reader, jobs chan []string) {
	for {
		rStr, err := fcsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		jobs <- rStr
	}
	defer close(jobs)
}

func waitAndClose(group *sync.WaitGroup, res chan *character.Characters) {
	group.Wait()
	close(res)
}

func addToRes(res chan *character.Characters, rs []*character.Characters) []*character.Characters {
	for r := range res {
		rs = append(rs, r)
	}
	return rs
}

func parseStruct(data []string) *character.Characters {
	id, _ := strconv.ParseInt(data[0], 10, 32)
	name := data[1]
	status := data[2]
	species := data[3]
	types := data[4]
	gender := data[5]
	return &character.Characters{
		ID:      int(id),
		Name:    name,
		Status:  status,
		Species: species,
		Type:    types,
		Gender:  gender,
	}
}

//NewCharacterRepository constructor of CharacterRepositoryStruct
func NewCharacterRepository() *CharacterRepositoryStruct {
	return &CharacterRepositoryStruct{}
}
