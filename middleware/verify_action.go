package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func VerifyCanProcessTradingProjectMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 1 || value == 2 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessBPOProjectMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 1 || value == 3 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessSPPHMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 5 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessSPHMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 6 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessNegotiationMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 7 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessSPKMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 8 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessBASTMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 9 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessPaymentMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 10 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func VerifyCanProcessRAPMiddleware(next http.Handler) http.Handler {
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

		permitted := false
		for _, value := range permissions {
			if value == 4 || value == 11 {
				permitted = true
			}
		}

		if !permitted {
			http.Error(writer, "Forbidden", 403)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
