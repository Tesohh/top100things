package handler

import (
	"net/http"

	"github.com/Tesohh/top100things/model"
	"github.com/Tesohh/top100things/render"
	"github.com/Tesohh/top100things/sb"
)

func Things(w http.ResponseWriter, r *http.Request) {
	var things []model.Thing
	sb.SB.DB.From("things").Select("*").Execute(&things)

	render.Template(w, "things", map[string]any{"things": things})
}
