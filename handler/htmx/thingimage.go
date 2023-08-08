package htmx

import (
	"fmt"
	"net/http"

	"github.com/Tesohh/top100things/sb"
)

var defaultImage = `<img id="imagepreview" src="https://cdn-icons-png.flaticon.com/512/17/17671.png" width="100" height="100">`

func ThingImage(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("title")
	if t == "" {
		fmt.Fprint(w, defaultImage)
	} else {
		var res []map[string]string
		sb.SB.DB.From("things").Select("image_url").Limit(1).Ilike("name", "%%"+t+"%%").Execute(&res)
		if len(res) == 0 {
			fmt.Fprint(w, defaultImage)
			return
		}
		img, ok := res[0]["image_url"]
		if !ok {
			fmt.Fprint(w, defaultImage)
			return
		}
		fmt.Fprintf(w, `<img id="imagepreview" src="%v">`, img)
	}
}
