package entity

type Film struct {
	FilmID string  `db:"id" json:"-"`
	Name   string  `db:"name" json:"name"`
	Year   string  `db:"year" json:"year"`
	Rating float64 `db:"rating" json:"ratting"`
}
