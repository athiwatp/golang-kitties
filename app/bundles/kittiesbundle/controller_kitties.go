package kittiesbundle

import (
	"net/http"

	"github.com/laibulle/kitties/app/core"
)

// KittiesController struct
type KittiesController struct {
	core.Controller
	km KittiesMapper
}

// NewKittiesController instance
func NewKittiesController(km KittiesMapper) *KittiesController {
	return &KittiesController{
		Controller: core.Controller{},
		km:         km,
	}
}

// Index func return all kitties in database
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	k, err := c.km.FindAll()

	if c.HandleError(err, w) {
		return
	}

	c.SendJSON(w, &k, http.StatusOK)
}
