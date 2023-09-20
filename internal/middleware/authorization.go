package middleware

import (
	"errors"
	"fmt"
	"net/http"
)

var UnAuthorizedError = errors.New(fmt.Sprintf("Invalid Id or Token."))

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
