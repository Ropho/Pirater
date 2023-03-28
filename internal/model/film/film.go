package model

type Film struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	PicUrl      string   `json:"url"`
	Hash        uint32   `json:"hash"`
	DescPath    string   `json:"description,omitempty"`
	FilmPath    string   `json:"film,omitempty"`
	TrailerPath string   `json:"trailer,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Rights      []string `json:"rights,omitempty"`
	Rating      int      `json:"rating"`
	Timestamp   string   `json:"timestamp"`
}
