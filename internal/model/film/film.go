package model

type Film struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PicUrl      string `json:"url"`
	Description string `json:"description,omitempty"`
	FilmPath    string `json:"-"`
	Category    string `json:"category,omitempty"`
	Rights      string `json:"-"`
}
