package handler

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/model"
	"github.com/Tesohh/top100things/render"
	"github.com/Tesohh/top100things/sb"
	"github.com/gorilla/mux"
)

func Things(w http.ResponseWriter, r *http.Request) {
	userid, ok := mux.Vars(r)["userid"]
	if !ok {
		fmt.Fprint(w, "You need to provide a user id")
		return
	}
	var rankings []model.Ranking
	sb.SB.DB.From("rankings").Select("*", "thing_id(*)").Eq("user_id", userid).Execute(&rankings)

	render.Template(w, "things", map[string]any{"rankings": rankings, "userid": userid})
}
