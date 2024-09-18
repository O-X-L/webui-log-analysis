package api

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, []string{"abc", "def"})
}
