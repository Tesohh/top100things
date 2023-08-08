package htmx

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HTMX is working")
}
