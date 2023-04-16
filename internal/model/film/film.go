package model

type Film struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Hash        uint32   `json:"hash"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	VideoUrl    string   `json:"video_url"`
	HeaderUrl   string   `json:"header_url"`
	AfishaUrl   string   `json:"afisha_url"`
	//CadreUrl []string
}
