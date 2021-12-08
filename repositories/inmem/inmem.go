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
		return characterSlice[i].ID <= characterSlice[j].ID
	})
	c.characters = characterSlice
	return nil
}

func (c *CharacterRepositoryStruct) WriteCharactersOnCsv(characters *[]character.Characters) error {
	csvFile, err := os.Create("../../sample-data/charactersWrite.csv")
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
	csvFile, err := os.Open("./sample-data/characters.csv")
	if err != nil {
		panic("asd")
	}

	fcsv := csv.NewReader(csvFile)
	rs := make([]*character.Characters, 0)
	numWps := math.Floor(float64(filter.Items / filter.ItemsPerWorker))
	jobs := make(chan []string)
	var res chan *character.Characters
	if filter.Items != -1 && filter.Items != 0 {
		res = make(chan *character.Characters, filter.Items)
	} else {
		res = make(chan *character.Characters)

	}
	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- *character.Characters, id int, numOfItems int64) {
		for i := 0; i < int(numOfItems); i++ {
			select {
			case job, ok := <-jobs:
				if !ok {
					return
				}
				newCharacter := parseStruct(job)
				if filter.TypeFilter != "" {
					if filter.TypeFilter == "odd" {
						if odd := newCharacter.ID % 2; odd == 0 {
							results <- newCharacter
						}
					} else {
						if even := newCharacter.ID % 2; even != 0 {
							results <- newCharacter
						}
					}
				} else {
					results <- parseStruct(job)
				}
			}
		}
	}

	// init workers
	for w := 0; w < int(numWps); w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res, w, filter.ItemsPerWorker)
		}()
	}

	go func() {
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
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		rs = append(rs, r)
	}

	fmt.Println("Count Concu ", len(rs))
	return rs, nil
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
