package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/kevin07696/ecommerce/templates/partials"
)

func HandleCreateAccount(authAPI services.API) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		input := struct {
			Username string
			Email    string
			Role     string
			OTP      string
		}{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Role:     r.FormValue("role"),
			OTP:      r.FormValue("otp"),
		}

		userResp, err := authAPI.RegisterUser(ctx, services.RegisterUserReq{Email: input.Email, Username: input.Username, Role: input.Role, OTP: input.OTP})
		if err != nil {
			// errorHandler(w, err)
			partials.Response(err.Error()).Render(ctx, w)
			return
		}

		authAPI.CreateSession(ctx, w, r, services.CreateSessionReq(userResp))
		w.WriteHeader(http.StatusOK)
	}
}
