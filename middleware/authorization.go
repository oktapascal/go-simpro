package middleware

import (
	"net/http"
	"strings"
)

func AuthorizationCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")
		if authorization == "" {
			http.Error(writer, http.StatusText(401), 401)
			return
		}

		if !strings.Contains(authorization, "Bearer") {
			http.Error(writer, http.StatusText(400), 400)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
