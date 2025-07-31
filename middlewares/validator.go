package middlewares

import (
	"AuthApp/dto"
	"AuthApp/util"
	"context"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO

		// read and decode json body into payload
		err := util.ReadJsonBody(r, &payload)
		if err != nil {
			util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error reading json body", err)
			return
		}

		// validate the payload using the validator instance
		validationErr := util.Validator.Struct(payload)
		if validationErr != nil {
			util.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation failed", validationErr)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx)) // call the next handler in the chain
	})
}

func UserSignupRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload *dto.SignupUserRequestDTO

		// read and decode json body into payload
		err := util.ReadJsonBody(r, &payload)
		if err != nil {
			util.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error reading json body", err)
			return
		}

		// validate the payload using the validator instance
		validationErr := util.Validator.Struct(payload)
		if validationErr != nil {
			util.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation failed", validationErr)
			return
		}

		// fmt.Println("payload validated", payload)

		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx)) // call the next handler in the chain
	})
}
