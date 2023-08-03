package handler

import (
	"net/http"

	"github.com/Tesohh/top100things/model"
	"github.com/Tesohh/top100things/sb"
	"github.com/Tesohh/top100things/template"
)

func Things(w http.ResponseWriter, r *http.Request) {
	var things []model.Thing
	sb.SB.DB.From("things").Select("*").Execute(&things)

	template.Parse(w, "things", things)
}
