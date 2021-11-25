package characters

type Characters struct {
	Id      int    `json:"id" csv:"id"`
	Name    string `json:"name" csv:"name"`
	Status  string `json:"status" csv:"status"`
	Species string `json:"species" csv:"species"`
	Type    string `json:"type" csv:"type"`
	Gender  string `json:"gender" csv:"gender"`
}

type CharactersRepository interface {
	FetchCharacters() ([]*Characters, error)
	FetchCharacterById(id int) (*Characters, error)
}
