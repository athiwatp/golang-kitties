package kittiesbundle

// KittiesMapper define the base contract for mapper
type KittiesMapper interface {
	FindAll() ([]Kitty, error)
}
