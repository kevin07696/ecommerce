package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/kevin07696/ecommerce/templates/partials"
)

func HandleSendingRegisterOTP(authAPI services.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		input := struct {
			Email string
		}{
			Email: r.FormValue("email"),
		}

		err := authAPI.SendRegisterOTP(ctx, services.SendRegisterOTPReq{Email: input.Email})
		if err != nil {
			// errorHandler(w, err)
			partials.Response(err.Error()).Render(ctx, w)
		}

		w.WriteHeader(http.StatusOK)
	}
}
