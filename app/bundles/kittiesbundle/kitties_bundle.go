package kittiesbundle

import (
	"github.com/jinzhu/gorm"
	"github.com/laibulle/kitties/app/core"
)

// KittiesBundle handle kitties resources
type KittiesBundle struct {
	routes []core.Route
}

// NewKittiesBundle instance
func NewKittiesBundle(db *gorm.DB) core.Bundle {
	km := NewKittiesSQLMapper(db)
	kc := NewKittiesController(km)

	r := []core.Route{
		core.Route{
			Method:  core.GET,
			Path:    "/kitties",
			Handler: kc.Index,
		},
	}

	return &KittiesBundle{
		routes: r,
	}
}

// GetRoutes implement interface core.Bundle
func (b *KittiesBundle) GetRoutes() []core.Route {
	return b.routes
}
