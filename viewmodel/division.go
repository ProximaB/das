package viewmodel

type SearchDivisionViewModel struct {
	DivisionID   int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type DivisionViewModel struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Federation int    `json:"federation"`
}
