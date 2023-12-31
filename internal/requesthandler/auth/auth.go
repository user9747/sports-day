package auth_req

import (
	"context"
	"encoding/json"
	"net/http"
	"sports-day/internal/errorhandler"
	"sports-day/internal/middlewares"
	"sports-day/internal/usecases/auth"
)

var authService = auth.GetAuthService()

// Login controller returns token after checking
func Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		var reqObj UserLoginReq
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&reqObj)
		if err != nil {
			panic(err.Error())
		}

		token, err := authService.Login(r.Context(), reqObj.UserName)

		if err != nil {
			panic(err.Error())
		}

		var response = map[string]interface{}{
			"token": token,
		}
		ctx := context.WithValue(r.Context(), "resData", response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logout logsout the current logged in user
func Logout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)
		token := r.Header.Get(middlewares.HeaderToken)
		err := authService.Logout(r.Context(), token)
		if err != nil {
			panic(err.Error())
		}

		ctx := context.WithValue(r.Context(), "resData", "success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
