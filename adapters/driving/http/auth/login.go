package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/kevin07696/ecommerce/templates/partials"
)

func HandleLogin(userAPI services.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		input := struct {
			UserId string
			OTP   string
		}{
			UserId: r.FormValue("user_id"),
			OTP:   r.FormValue("otp"),
		}

		userResp, err := userAPI.LoginUser(ctx, services.LoginUserReq{UserId: input.UserId, OTP: input.OTP})
		if err != nil {
			// errorHandler(w, err)
			partials.Response(err.Error()).Render(ctx, w)
			return
		}

		userAPI.CreateSession(ctx, w, r, services.CreateSessionReq{Username: userResp.Username})
		w.WriteHeader(http.StatusOK)
	}
}
