package domain

//var (
//	ErrMusicNotFound = errors.New("music not found")
//)

type ListMusic []Music

type Music struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Performer   string `json:"performer"`
	RealiseYear int    `json:"realise_year"`
	Genre       string `json:"genre"`
}

type UpdateMusicInput struct {
	Name        *string `json:"name"`
	Performer   *string `json:"performer"`
	RealiseYear *int    `json:"realise_year"`
	Genre       *string `json:"genre"`
}
