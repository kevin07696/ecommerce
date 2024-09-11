package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/kevin07696/ecommerce/templates/partials"
)

func HandleValidateUsername(userAPI services.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		input := struct {
			Username string
		}{
			Username: r.URL.Query().Get("username"),
		}

		err := userAPI.ValidateUsername(ctx, services.ValidateUsernameReq{Username: input.Username})
		if err != nil {
			// errorHandler(w, err)
			partials.Response(err.Error()).Render(ctx, w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
