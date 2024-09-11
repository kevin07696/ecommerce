package web

import (
	"net/http"

	"github.com/a-h/templ"
)

func HandleComponents(c templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Render(r.Context(), w)
	}
}
