package characters

//Characters - it's a Struct to manage the basic info of a Character
type Characters struct {
	ID      int    `json:"id" csv:"id"`
	Name    string `json:"name" csv:"name"`
	Status  string `json:"status" csv:"status"`
	Species string `json:"species" csv:"species"`
	Type    string `json:"type" csv:"type"`
	Gender  string `json:"gender" csv:"gender"`
}

//CharactersRepository - Interface that contains two functions who are the ones that retrieve info of the CSVFiles
type CharactersRepository interface {
	FetchCharacters() ([]*Characters, error)
	FetchCharacterById(id int) (*Characters, error)
}
