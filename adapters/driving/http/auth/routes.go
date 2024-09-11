package auth

import (
	"net/http"

	web "github.com/kevin07696/ecommerce/adapters/driving/http"
	"github.com/kevin07696/ecommerce/templates/pages"
)

func (h *Handler) initAppRoutes() {
	h.router.HandleFunc("POST /api/register", HandleCreateAccount(h.userAPI))
	h.router.HandleFunc("POST /api/login", HandleLogin(h.userAPI))
	h.router.HandleFunc("POST /api/login-otp", HandleSendLoginOTP(h.userAPI))
	h.router.HandleFunc("POST /api/valid-email", HandleSendingRegisterOTP(h.userAPI))
	h.router.HandleFunc("GET /api/valid-username", HandleValidateUsername(h.userAPI))

	h.router.HandleFunc("GET /register", web.HandleComponents(pages.RegisterForm()))
	h.router.HandleFunc("GET /login", web.HandleComponents(pages.LoginForm()))
	h.router.HandleFunc("GET /test", web.HandleComponents(pages.TestComponents()))

	h.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}