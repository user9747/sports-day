package admin_req

import (
	"context"
	"encoding/json"
	"net/http"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/users"
	"sports-day/internal/utils"
)

var userService = users.GetUserService()

func CreateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)

		var reqObj UserReq
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&reqObj)

		if err != nil {
			panic("invalid request")
		}

		userID, err := userService.CreateUser(r.Context(), reqObj.UserName, reqObj.Role)
		if err != nil {
			if utils.IsUniqueConstraintViolation(err) {
				panic("username already exists")
			}
			panic(err.Error())
		}

		var response = map[string]interface{}{
			"userID": userID,
		}
		ctx := context.WithValue(r.Context(), "resData", response)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
