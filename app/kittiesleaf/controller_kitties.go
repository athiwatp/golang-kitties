package kittiesleaf

import (
	"net/http"
	"time"

	"github.com/laibulle/kitties/app/common"
)

// KittiesController struct
type KittiesController struct {
	common.Controller
}

// Index func return all kitties in database
func (c *KittiesController) Index(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		[]*Kitty{NewKitty("Gaspart", "British", time.Now())},
		http.StatusOK,
	)
}
