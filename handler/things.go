package handler

import (
	"net/http"

	"github.com/Tesohh/top100things/model"
	"github.com/Tesohh/top100things/render"
	"github.com/Tesohh/top100things/sb"
)

func Things(w http.ResponseWriter, r *http.Request) {
	var rankings []model.Ranking
	sb.SB.DB.From("rankings").Select("*", "thing_id(*)").Execute(&rankings)

	render.Template(w, "things", map[string]any{"rankings": rankings})
}
