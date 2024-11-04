package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func VerifyRootUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userInfo := request.Context().Value("claims").(jwt.MapClaims)

		var permissions []int
		for index, value := range userInfo["permissions"].([]any) {
			permission, ok := value.(float64)
			if !ok {
				panic(fmt.Sprintf("Invalid type in permissions array at index %d", index))
			}

			permissions = append(permissions, int(permission))
		}

		IDRole := userInfo["aud"]
		if IDRole != 1 {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
