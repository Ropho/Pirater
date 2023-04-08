package model

type Film struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	PicUrl      string   `json:"pic_url"`
	Hash        uint32   `json:"hash"`
	Description string   `json:"description,omitempty"`
	FilmUrl     string   `json:"film_url,omitempty"`
	TrailerUrl  string   `json:"trailer_url,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Rights      []string `json:"rights,omitempty"`
	Rating      int      `json:"rating"`
	Timestamp   string   `json:"timestamp"`
}
