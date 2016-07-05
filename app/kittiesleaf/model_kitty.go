package kittiesleaf

import "time"

// Kitty struct
type Kitty struct {
	Name      string    `json:"name"`
	Breed     string    `json:"breed"`
	BirthDate time.Time `json:"birthDate"`
}

// NewKitty create a new Kitty
func NewKitty(name string, breed string, birthDate time.Time) *Kitty {
	return &Kitty{
		Name:      name,
		Breed:     breed,
		BirthDate: birthDate,
	}
}
