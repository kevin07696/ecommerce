package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/kevin07696/ecommerce/templates/partials"
)

func HandleSendLoginOTP(userAPI services.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		input := struct {
			UserId string
		}{
			UserId: r.FormValue("user_id"),
		}

		err := userAPI.SendLoginOTP(ctx, services.SendLoginOTPReq{UserId: input.UserId})
		if err != nil {
			// errorHandler(w, err)
			partials.Response(err.Error()).Render(ctx, w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
