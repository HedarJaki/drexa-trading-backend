package adapter

import (
	"drexa/internal/shared/utils"
	"net/http"
)

func HandleRegister() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			type request struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			body, err := utils.Decode[request](r)
			if err != nil || body.Email == "" || body.Password == "" {
				utils.Encode(w, r, http.StatusConflict, map[string]string{
					"error": "invalid request body",
				})
				return
			}

			// if db.checkEmail(body.Email)  {
			// 	utils.Encode(w, r, http.StatusBadRequest, map[string]string{
			// 		"error": "email already exist",
			// 	})
			// 	return
			// }

			// auth.createAccount(body.Email, body.Password)

			utils.Encode(w, r, http.StatusOK, map[string]string{
				"otp": "4040",
			})
		},
	)
}
