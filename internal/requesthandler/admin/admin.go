package admin_req

import (
	"context"
	"encoding/json"
	"net/http"
	"sports-day/internal/entity"
	"sports-day/internal/errorhandler"
	"sports-day/internal/usecases/users"
	"sports-day/internal/utils"
)

var userService = users.GetUserService()

// CreateUser creates user from userName and role
func CreateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorhandler.Recovery(w, r, http.StatusBadRequest)

		loggedInUser := utils.GetLoggedInUser(r.Context())

		if loggedInUser == nil {
			panic("invalid user")
		}

		if loggedInUser.Role != entity.UserRoleAdmin {
			errorhandler.CustomError(w, http.StatusForbidden, "not admin")
			return
		}

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
