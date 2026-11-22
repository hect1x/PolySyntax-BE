package domain

type Snippet struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Objective struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Snippets []Snippet `json:"snippets"`
}
