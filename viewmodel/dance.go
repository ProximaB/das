package viewmodel

type Dance struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	StyleID      int    `json:"style"`
	Abbreviation string `json:"abbreviation"`
}
